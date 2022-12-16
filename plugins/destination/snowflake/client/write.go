package client

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/cloudquery/plugin-sdk/schema"
)

const (
	createOrReplaceFileFormat = `create or replace file format cq_plugin_json_format type = 'JSON'`
	createOrReplaceStage      = `create or replace stage cq_plugin_stage file_format = cq_plugin_json_format;`
	putFileIntoStage          = `put file://%s @cq_plugin_stage auto_compress=true`
	copyIntoTable             = `copy into %s from @cq_plugin_stage/%s file_format = (format_name = cq_plugin_json_format) match_by_column_name = case_insensitive`
)

func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, resources [][]interface{}) error {
	f, err := os.CreateTemp(os.TempDir(), table.Name+".json.*")
	if err != nil {
		return err
	}
	defer func() {
		f.Close()
		os.Remove(f.Name())
	}()

	for _, r := range resources {
		jsonObj := make(map[string]interface{}, len(table.Columns))
		for i := range r {
			jsonObj[table.Columns[i].Name] = r[i]
		}

		b, err := json.Marshal(jsonObj)
		if err != nil {
			return err
		}
		b = append(b, '\n')
		if _, err := f.Write(b); err != nil {
			return err
		}
	}

	if err := f.Close(); err != nil {
		return err
	}
	sql := fmt.Sprintf(putFileIntoStage, f.Name())
	if _, err := c.db.ExecContext(ctx, sql); err != nil {
		return fmt.Errorf("failed to put file into stage %s: %w", sql, err)
	}
	if err := os.Remove(f.Name()); err != nil {
		return err
	}
	sql = fmt.Sprintf(copyIntoTable, table.Name, path.Base(f.Name()))
	if _, err := c.db.ExecContext(ctx, sql); err != nil {
		return fmt.Errorf("failed to copy file into table %s: %w", sql, err)
	}
	return nil
}

func (c *Client) PreWrite(ctx context.Context, tables schema.Tables, source string, syncTime time.Time) error {
	if _, err := c.db.ExecContext(ctx, createOrReplaceFileFormat); err != nil {
		return fmt.Errorf("failed to create file format %s: %w", createOrReplaceFileFormat, err)
	}

	if _, err := c.db.ExecContext(ctx, createOrReplaceStage); err != nil {
		return fmt.Errorf("failed to create stage %s: %w", createOrReplaceStage, err)
	}
	return nil
}
