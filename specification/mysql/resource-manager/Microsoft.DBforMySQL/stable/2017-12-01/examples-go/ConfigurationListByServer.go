package armmysql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ConfigurationListByServer.json
func ExampleConfigurationsClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConfigurationsClient().NewListByServerPager("testrg", "mysqltestsvc1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ConfigurationListResult = armmysql.ConfigurationListResult{
		// 	Value: []*armmysql.Configuration{
		// 		{
		// 			Name: to.Ptr("event_scheduler"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/event_scheduler"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("Indicates the status of the Event Scheduler."),
		// 				AllowedValues: to.Ptr("ON,OFF,DISABLED"),
		// 				DataType: to.Ptr("Enumeration"),
		// 				DefaultValue: to.Ptr("OFF"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("OFF"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("div_precision_increment"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/div_precision_increment"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("Number of digits by which to increase the scale of the result of division operations."),
		// 				AllowedValues: to.Ptr("0-30"),
		// 				DataType: to.Ptr("Integer"),
		// 				DefaultValue: to.Ptr("4"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("4"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("group_concat_max_len"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/group_concat_max_len"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("Maximum allowed result length in bytes for the GROUP_CONCAT()."),
		// 				AllowedValues: to.Ptr("4-16777216"),
		// 				DataType: to.Ptr("Integer"),
		// 				DefaultValue: to.Ptr("1024"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("1024"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("innodb_adaptive_hash_index"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/innodb_adaptive_hash_index"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("Whether innodb adaptive hash indexes are enabled or disabled."),
		// 				AllowedValues: to.Ptr("ON,OFF"),
		// 				DataType: to.Ptr("Enumeration"),
		// 				DefaultValue: to.Ptr("ON"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("ON"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("innodb_lock_wait_timeout"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/innodb_lock_wait_timeout"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("The length of time in seconds an InnoDB transaction waits for a row lock before giving up."),
		// 				AllowedValues: to.Ptr("1-3600"),
		// 				DataType: to.Ptr("Integer"),
		// 				DefaultValue: to.Ptr("50"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("50"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("interactive_timeout"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/interactive_timeout"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("Number of seconds the server waits for activity on an interactive connection before closing it."),
		// 				AllowedValues: to.Ptr("10-1800"),
		// 				DataType: to.Ptr("Integer"),
		// 				DefaultValue: to.Ptr("1800"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("1800"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("log_queries_not_using_indexes"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/log_queries_not_using_indexes"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("Logs queries that are expected to retrieve all rows to slow query log."),
		// 				AllowedValues: to.Ptr("ON,OFF"),
		// 				DataType: to.Ptr("Enumeration"),
		// 				DefaultValue: to.Ptr("OFF"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("OFF"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("log_throttle_queries_not_using_indexes"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/log_throttle_queries_not_using_indexes"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("Limits the number of such queries per minute that can be written to the slow query log."),
		// 				AllowedValues: to.Ptr("0-4294967295"),
		// 				DataType: to.Ptr("Integer"),
		// 				DefaultValue: to.Ptr("0"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("0"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("log_slow_admin_statements"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/log_slow_admin_statements"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("Include slow administrative statements in the statements written to the slow query log."),
		// 				AllowedValues: to.Ptr("ON,OFF"),
		// 				DataType: to.Ptr("Enumeration"),
		// 				DefaultValue: to.Ptr("OFF"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("OFF"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("log_slow_slave_statements"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/log_slow_slave_statements"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("When the slow query log is enabled, this variable enables logging for queries that have taken more than long_query_time seconds to execute on the slave."),
		// 				AllowedValues: to.Ptr("ON,OFF"),
		// 				DataType: to.Ptr("Enumeration"),
		// 				DefaultValue: to.Ptr("OFF"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("OFF"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("log_bin_trust_function_creators"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/log_bin_trust_function_creators"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("This variable applies when binary logging is enabled. It controls whether stored function creators can be trusted not to create stored functions that will cause unsafe events to be written to the binary log."),
		// 				AllowedValues: to.Ptr("ON,OFF"),
		// 				DataType: to.Ptr("Enumeration"),
		// 				DefaultValue: to.Ptr("OFF"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("OFF"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("long_query_time"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/long_query_time"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("If a query takes longer than this many seconds, the server increments the Slow_queries status variable."),
		// 				AllowedValues: to.Ptr("0-1E+100"),
		// 				DataType: to.Ptr("Numeric"),
		// 				DefaultValue: to.Ptr("10"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("10"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("min_examined_row_limit"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/min_examined_row_limit"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("Can be used to cause queries which examine fewer than the stated number of rows not to be logged."),
		// 				AllowedValues: to.Ptr("0-18446744073709551615"),
		// 				DataType: to.Ptr("Integer"),
		// 				DefaultValue: to.Ptr("0"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("0"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("slow_query_log"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/slow_query_log"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("Enable or disable the slow query log"),
		// 				AllowedValues: to.Ptr("ON,OFF"),
		// 				DataType: to.Ptr("Enumeration"),
		// 				DefaultValue: to.Ptr("OFF"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("OFF"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sql_mode"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/sql_mode"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("The current server SQL mode."),
		// 				AllowedValues: to.Ptr(",ALLOW_INVALID_DATES,ANSI_QUOTES,ERROR_FOR_DIVISION_BY_ZERO,HIGH_NOT_PRECEDENCE,IGNORE_SPACE,NO_AUTO_CREATE_USER,NO_AUTO_VALUE_ON_ZERO,NO_BACKSLASH_ESCAPES,NO_DIR_IN_CREATE,NO_ENGINE_SUBSTITUTION,NO_FIELD_OPTIONS,NO_KEY_OPTIONS,NO_TABLE_OPTIONS,NO_UNSIGNED_SUBTRACTION,NO_ZERO_DATE,NO_ZERO_IN_DATE,ONLY_FULL_GROUP_BY,PAD_CHAR_TO_FULL_LENGTH,PIPES_AS_CONCAT,REAL_AS_FLOAT,STRICT_ALL_TABLES,STRICT_TRANS_TABLES"),
		// 				DataType: to.Ptr("Set"),
		// 				DefaultValue: to.Ptr(""),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr(""),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("wait_timeout"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/wait_timeout"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("The number of seconds the server waits for activity on a noninteractive connection before closing it."),
		// 				AllowedValues: to.Ptr("60-86400"),
		// 				DataType: to.Ptr("Integer"),
		// 				DefaultValue: to.Ptr("120"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("120"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("net_read_timeout"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/net_read_timeout"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("The number of seconds the server waits for network reading action, especially for LOAD DATA LOCAL FILE."),
		// 				AllowedValues: to.Ptr("10-3600"),
		// 				DataType: to.Ptr("Integer"),
		// 				DefaultValue: to.Ptr("120"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("120"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("net_write_timeout"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/net_write_timeout"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("The number of seconds the server waits for network writing action, especially for LOAD DATA LOCAL FILE."),
		// 				AllowedValues: to.Ptr("10-3600"),
		// 				DataType: to.Ptr("Integer"),
		// 				DefaultValue: to.Ptr("240"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("240"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("server_id"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/server_id"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("The server ID, used in replication to give each master and slave a unique identity."),
		// 				AllowedValues: to.Ptr("1000-4294967295"),
		// 				DataType: to.Ptr("Integer"),
		// 				DefaultValue: to.Ptr("1000"),
		// 				Source: to.Ptr("user-override"),
		// 				Value: to.Ptr("1381286943"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("max_allowed_packet"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/max_allowed_packet"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("The maximum size of one packet or any generated/intermediate string, or any parameter sent by the mysql_stmt_send_long_data() C API function."),
		// 				AllowedValues: to.Ptr("1024-1073741824"),
		// 				DataType: to.Ptr("Integer"),
		// 				DefaultValue: to.Ptr("536870912"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("536870912"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("slave_net_timeout"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/slave_net_timeout"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("The number of seconds to wait for more data from the master before the slave considers the connection broken, aborts the read, and tries to reconnect."),
		// 				AllowedValues: to.Ptr("30-3600"),
		// 				DataType: to.Ptr("Integer"),
		// 				DefaultValue: to.Ptr("60"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("60"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("time_zone"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/time_zone"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("The server time zone"),
		// 				AllowedValues: to.Ptr("[+|-][0]{0,1}[0-9]:[0-5][0-9]|[+|-][1][0-2]:[0-5][0-9]|SYSTEM"),
		// 				DataType: to.Ptr("String"),
		// 				DefaultValue: to.Ptr("SYSTEM"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("SYSTEM"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("binlog_group_commit_sync_delay"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/binlog_group_commit_sync_delay"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("Controls how many microseconds the binary log commit waits before synchronizing the binary log file to disk."),
		// 				AllowedValues: to.Ptr("0,11-1000000"),
		// 				DataType: to.Ptr("Integer"),
		// 				DefaultValue: to.Ptr("1000"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("1000"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("binlog_group_commit_sync_no_delay_count"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/binlog_group_commit_sync_no_delay_count"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("The maximum number of transactions to wait for before aborting the current delay as specified by binlog-group-commit-sync-delay."),
		// 				AllowedValues: to.Ptr("0-1000000"),
		// 				DataType: to.Ptr("Integer"),
		// 				DefaultValue: to.Ptr("0"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("0"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("character_set_server"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/character_set_server"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("Use charset_name as the default server character set."),
		// 				AllowedValues: to.Ptr("BIG5,DEC8,CP850,HP8,KOI8R,LATIN1,LATIN2,SWE7,ASCII,UJIS,SJIS,HEBREW,TIS620,EUCKR,KOI8U,GB2312,GREEK,CP1250,GBK,LATIN5,ARMSCII8,UTF8,UCS2,CP866,KEYBCS2,MACCE,MACROMAN,CP852,LATIN7,UTF8MB4,CP1251,UTF16,CP1256,CP1257,UTF32,BINARY,GEOSTD8,CP932,EUCJPMS"),
		// 				DataType: to.Ptr("Enumeration"),
		// 				DefaultValue: to.Ptr("latin1"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("latin1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("join_buffer_size"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/join_buffer_size"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("The minimum size of the buffer that is used for plain index scans, range index scans, and joins that do not use indexes and thus perform full table scans."),
		// 				AllowedValues: to.Ptr("128-2097152"),
		// 				DataType: to.Ptr("Integer"),
		// 				DefaultValue: to.Ptr("262144"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("262144"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("table_open_cache"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/table_open_cache"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("The number of open tables for all threads."),
		// 				AllowedValues: to.Ptr("1-4000"),
		// 				DataType: to.Ptr("Integer"),
		// 				DefaultValue: to.Ptr("2000"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("2000"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("lower_case_table_names"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/lower_case_table_names"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("If set to 1, table names are stored in lowercase on disk and comparisons are not case sensitive. If set to 2, table names are stored as given but compared in lowercase."),
		// 				AllowedValues: to.Ptr("1,2"),
		// 				DataType: to.Ptr("Enumeration"),
		// 				DefaultValue: to.Ptr("1"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("slave_compressed_protocol"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/slave_compressed_protocol"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("This option places an upper limit on the total size in bytes of all relay logs on the slave."),
		// 				AllowedValues: to.Ptr("ON,OFF"),
		// 				DataType: to.Ptr("Enumeration"),
		// 				DefaultValue: to.Ptr("OFF"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("OFF"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("innodb_io_capacity"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/innodb_io_capacity"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("Sets an upper limit on I/O activity performed by InnoDB background tasks, such as flushing pages from the buffer pool and merging data from the change buffer."),
		// 				AllowedValues: to.Ptr("100-1500"),
		// 				DataType: to.Ptr("Integer"),
		// 				DefaultValue: to.Ptr("200"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("200"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("innodb_read_io_threads"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/innodb_read_io_threads"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("The number of I/O threads for read operations in InnoDB."),
		// 				AllowedValues: to.Ptr("1-64"),
		// 				DataType: to.Ptr("Integer"),
		// 				DefaultValue: to.Ptr("4"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("4"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("innodb_thread_concurrency"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/innodb_thread_concurrency"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("InnoDB tries to keep the number of operating system threads concurrently inside InnoDB less than or equal to the limit given by this variable."),
		// 				AllowedValues: to.Ptr("0-1000"),
		// 				DataType: to.Ptr("Integer"),
		// 				DefaultValue: to.Ptr("0"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("0"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("innodb_write_io_threads"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/innodb_write_io_threads"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("The number of I/O threads for write operations in InnoDB."),
		// 				AllowedValues: to.Ptr("1-64"),
		// 				DataType: to.Ptr("Integer"),
		// 				DefaultValue: to.Ptr("4"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("4"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("innodb_page_cleaners"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/innodb_page_cleaners"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("The number of page cleaner threads that flush dirty pages from buffer pool instances."),
		// 				AllowedValues: to.Ptr("1-64"),
		// 				DataType: to.Ptr("Integer"),
		// 				DefaultValue: to.Ptr("4"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("4"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("innodb_online_alter_log_max_size"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/innodb_online_alter_log_max_size"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("Specifies an upper limit on the size of the temporary log files used during online DDL operations for InnoDB tables."),
		// 				AllowedValues: to.Ptr("65536-2147483648"),
		// 				DataType: to.Ptr("Integer"),
		// 				DefaultValue: to.Ptr("134217728"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("134217728"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("init_connect"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/init_connect"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("A string to be executed by the server for each client that connects."),
		// 				AllowedValues: to.Ptr(""),
		// 				DataType: to.Ptr("String"),
		// 				DefaultValue: to.Ptr(""),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr(""),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("tx_isolation"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/tx_isolation"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("The default transaction isolation level."),
		// 				AllowedValues: to.Ptr("READ-UNCOMMITTED,READ-COMMITTED,REPEATABLE-READ,SERIALIZABLE"),
		// 				DataType: to.Ptr("Enumeration"),
		// 				DefaultValue: to.Ptr("REPEATABLE-READ"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("REPEATABLE-READ"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("eq_range_index_dive_limit"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/eq_range_index_dive_limit"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("This variable indicates the number of equality ranges in an equality comparison condition when the optimizer should switch from using index dives to index statistics in estimating the number of qualifying rows."),
		// 				AllowedValues: to.Ptr("0-4294967295"),
		// 				DataType: to.Ptr("Integer"),
		// 				DefaultValue: to.Ptr("200"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("200"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("innodb_old_blocks_pct"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/innodb_old_blocks_pct"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("Specifies the approximate percentage of the InnoDB buffer pool used for the old block sublist."),
		// 				AllowedValues: to.Ptr("5-95"),
		// 				DataType: to.Ptr("Integer"),
		// 				DefaultValue: to.Ptr("37"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("37"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("innodb_old_blocks_time"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/innodb_old_blocks_time"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("Non-zero values protect against the buffer pool being filled by data that is referenced only for a brief period, such as during a full table scan."),
		// 				AllowedValues: to.Ptr("0-4294967295"),
		// 				DataType: to.Ptr("Integer"),
		// 				DefaultValue: to.Ptr("1000"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("1000"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("innodb_read_ahead_threshold"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/innodb_read_ahead_threshold"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("Controls the sensitivity of linear read-ahead that InnoDB uses to prefetch pages into the buffer pool."),
		// 				AllowedValues: to.Ptr("0-64"),
		// 				DataType: to.Ptr("Integer"),
		// 				DefaultValue: to.Ptr("56"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("56"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("max_length_for_sort_data"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/max_length_for_sort_data"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("The cutoff on the size of index values that determines which filesort algorithm to use."),
		// 				AllowedValues: to.Ptr("4-8388608"),
		// 				DataType: to.Ptr("Integer"),
		// 				DefaultValue: to.Ptr("1024"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("1024"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("max_connect_errors"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/max_connect_errors"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("If more than this many successive connection requests from a host are interrupted without a successful connection, the server blocks that host from further connections."),
		// 				AllowedValues: to.Ptr("1-18446744073709551615"),
		// 				DataType: to.Ptr("Integer"),
		// 				DefaultValue: to.Ptr("100"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("100"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("innodb_thread_sleep_delay"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/innodb_thread_sleep_delay"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("Defines how long InnoDB threads sleep before joining the InnoDB queue, in microseconds."),
		// 				AllowedValues: to.Ptr("0-1000000"),
		// 				DataType: to.Ptr("Integer"),
		// 				DefaultValue: to.Ptr("10000"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("10000"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("innodb_file_format"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1/configurations/innodb_file_format"),
		// 			Properties: &armmysql.ConfigurationProperties{
		// 				Description: to.Ptr("Indicates the InnoDB file format for file-per-table tablespaces."),
		// 				AllowedValues: to.Ptr("Antelope,Barracuda"),
		// 				DataType: to.Ptr("Enumeration"),
		// 				DefaultValue: to.Ptr("Barracuda"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("Barracuda"),
		// 			},
		// 	}},
		// }
	}
}
