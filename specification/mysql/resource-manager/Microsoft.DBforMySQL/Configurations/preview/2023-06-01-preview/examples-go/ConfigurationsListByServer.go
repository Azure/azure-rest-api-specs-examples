package armmysqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8691fbfca8fcdc5a241a0b501c32fd4a76bb0cd/specification/mysql/resource-manager/Microsoft.DBforMySQL/Configurations/preview/2023-06-01-preview/examples/ConfigurationsListByServer.json
func ExampleConfigurationsClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConfigurationsClient().NewListByServerPager("testrg", "mysqltestserver", &armmysqlflexibleservers.ConfigurationsClientListByServerOptions{Tags: nil,
		Keyword:  nil,
		Page:     to.Ptr[int32](1),
		PageSize: to.Ptr[int32](8),
	})
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
		// page.ConfigurationListResult = armmysqlflexibleservers.ConfigurationListResult{
		// 	Value: []*armmysqlflexibleservers.Configuration{
		// 		{
		// 			Name: to.Ptr("archive"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/flexibleServers/mysqltestserver/configurations/archive"),
		// 			Properties: &armmysqlflexibleservers.ConfigurationProperties{
		// 				Description: to.Ptr("Tell the server to enable or disable archive engine."),
		// 				AllowedValues: to.Ptr("ON,OFF"),
		// 				DataType: to.Ptr("Enumeration"),
		// 				DefaultValue: to.Ptr("OFF"),
		// 				IsConfigPendingRestart: to.Ptr(armmysqlflexibleservers.IsConfigPendingRestartFalse),
		// 				IsDynamicConfig: to.Ptr(armmysqlflexibleservers.IsDynamicConfigFalse),
		// 				IsReadOnly: to.Ptr(armmysqlflexibleservers.IsReadOnlyTrue),
		// 				Source: to.Ptr(armmysqlflexibleservers.ConfigurationSourceSystemDefault),
		// 				Value: to.Ptr("OFF"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("audit_log_enabled"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/flexibleServers/mysqltestserver/configurations/audit_log_enabled"),
		// 			Properties: &armmysqlflexibleservers.ConfigurationProperties{
		// 				Description: to.Ptr("Allow to audit the log."),
		// 				AllowedValues: to.Ptr("ON,OFF"),
		// 				DataType: to.Ptr("Enumeration"),
		// 				DefaultValue: to.Ptr("OFF"),
		// 				IsConfigPendingRestart: to.Ptr(armmysqlflexibleservers.IsConfigPendingRestartFalse),
		// 				IsDynamicConfig: to.Ptr(armmysqlflexibleservers.IsDynamicConfigTrue),
		// 				IsReadOnly: to.Ptr(armmysqlflexibleservers.IsReadOnlyFalse),
		// 				Source: to.Ptr(armmysqlflexibleservers.ConfigurationSourceSystemDefault),
		// 				Value: to.Ptr("OFF"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("audit_log_events"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/flexibleServers/mysqltestserver/configurations/audit_log_events"),
		// 			Properties: &armmysqlflexibleservers.ConfigurationProperties{
		// 				Description: to.Ptr("Select the events to audit logs."),
		// 				AllowedValues: to.Ptr("DDL,DML_SELECT,DML_NONSELECT,DCL,ADMIN,DML,GENERAL,CONNECTION,TABLE_ACCESS"),
		// 				DataType: to.Ptr("Set"),
		// 				DefaultValue: to.Ptr("CONNECTION"),
		// 				IsConfigPendingRestart: to.Ptr(armmysqlflexibleservers.IsConfigPendingRestartFalse),
		// 				IsDynamicConfig: to.Ptr(armmysqlflexibleservers.IsDynamicConfigTrue),
		// 				IsReadOnly: to.Ptr(armmysqlflexibleservers.IsReadOnlyFalse),
		// 				Source: to.Ptr(armmysqlflexibleservers.ConfigurationSourceSystemDefault),
		// 				Value: to.Ptr("CONNECTION"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("audit_log_exclude_users"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/flexibleServers/mysqltestserver/configurations/audit_log_exclude_users"),
		// 			Properties: &armmysqlflexibleservers.ConfigurationProperties{
		// 				Description: to.Ptr("The comma-separated user list whose commands will not be in the audit logs."),
		// 				AllowedValues: to.Ptr(""),
		// 				DataType: to.Ptr("String"),
		// 				DefaultValue: to.Ptr("azure_superuser"),
		// 				IsConfigPendingRestart: to.Ptr(armmysqlflexibleservers.IsConfigPendingRestartFalse),
		// 				IsDynamicConfig: to.Ptr(armmysqlflexibleservers.IsDynamicConfigTrue),
		// 				IsReadOnly: to.Ptr(armmysqlflexibleservers.IsReadOnlyFalse),
		// 				Source: to.Ptr(armmysqlflexibleservers.ConfigurationSourceSystemDefault),
		// 				Value: to.Ptr("azure_superuser"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("audit_log_include_users"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/flexibleServers/mysqltestserver/configurations/audit_log_include_users"),
		// 			Properties: &armmysqlflexibleservers.ConfigurationProperties{
		// 				Description: to.Ptr("The comma-separated user list whose commands will be in the audit logs. It takes higher priority if the same user name is found in audit_log_exclude_users."),
		// 				AllowedValues: to.Ptr(""),
		// 				DataType: to.Ptr("String"),
		// 				DefaultValue: to.Ptr(""),
		// 				IsConfigPendingRestart: to.Ptr(armmysqlflexibleservers.IsConfigPendingRestartFalse),
		// 				IsDynamicConfig: to.Ptr(armmysqlflexibleservers.IsDynamicConfigTrue),
		// 				IsReadOnly: to.Ptr(armmysqlflexibleservers.IsReadOnlyFalse),
		// 				Source: to.Ptr(armmysqlflexibleservers.ConfigurationSourceSystemDefault),
		// 				Value: to.Ptr(""),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("audit_slow_log_enabled"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/flexibleServers/mysqltestserver/configurations/audit_slow_log_enabled"),
		// 			Properties: &armmysqlflexibleservers.ConfigurationProperties{
		// 				Description: to.Ptr("Allow to audit the slow log."),
		// 				AllowedValues: to.Ptr("ON,OFF"),
		// 				DataType: to.Ptr("Enumeration"),
		// 				DefaultValue: to.Ptr("ON"),
		// 				IsConfigPendingRestart: to.Ptr(armmysqlflexibleservers.IsConfigPendingRestartFalse),
		// 				IsDynamicConfig: to.Ptr(armmysqlflexibleservers.IsDynamicConfigTrue),
		// 				IsReadOnly: to.Ptr(armmysqlflexibleservers.IsReadOnlyTrue),
		// 				Source: to.Ptr(armmysqlflexibleservers.ConfigurationSourceSystemDefault),
		// 				Value: to.Ptr("ON"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("auto_generate_certs"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/flexibleServers/mysqltestserver/configurations/auto_generate_certs"),
		// 			Properties: &armmysqlflexibleservers.ConfigurationProperties{
		// 				Description: to.Ptr("Controls whether the server autogenerates SSL key and certificate files in the data directory, if they do not already exist."),
		// 				AllowedValues: to.Ptr("ON,OFF"),
		// 				DataType: to.Ptr("Enumeration"),
		// 				DefaultValue: to.Ptr("OFF"),
		// 				IsConfigPendingRestart: to.Ptr(armmysqlflexibleservers.IsConfigPendingRestartFalse),
		// 				IsDynamicConfig: to.Ptr(armmysqlflexibleservers.IsDynamicConfigFalse),
		// 				IsReadOnly: to.Ptr(armmysqlflexibleservers.IsReadOnlyTrue),
		// 				Source: to.Ptr(armmysqlflexibleservers.ConfigurationSourceSystemDefault),
		// 				Value: to.Ptr("OFF"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("auto_increment_increment"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/flexibleServers/mysqltestserver/configurations/auto_increment_increment"),
		// 			Properties: &armmysqlflexibleservers.ConfigurationProperties{
		// 				Description: to.Ptr("The auto_increment_increment is intended for use with source-to-source replication, and can be used to control the operation of AUTO_INCREMENT columns."),
		// 				AllowedValues: to.Ptr("1-65535"),
		// 				DataType: to.Ptr("Integer"),
		// 				DefaultValue: to.Ptr("1"),
		// 				IsConfigPendingRestart: to.Ptr(armmysqlflexibleservers.IsConfigPendingRestartFalse),
		// 				IsDynamicConfig: to.Ptr(armmysqlflexibleservers.IsDynamicConfigTrue),
		// 				IsReadOnly: to.Ptr(armmysqlflexibleservers.IsReadOnlyFalse),
		// 				Source: to.Ptr(armmysqlflexibleservers.ConfigurationSourceSystemDefault),
		// 				Value: to.Ptr("1"),
		// 			},
		// 	}},
		// }
	}
}
