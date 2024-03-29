package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getAgentRegistration.json
func ExampleAgentRegistrationInformationClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAgentRegistrationInformationClient().Get(ctx, "rg", "myAutomationAccount18", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AgentRegistration = armautomation.AgentRegistration{
	// 	DscMetaConfiguration: to.Ptr("\r\n	instance of MSFT_WebDownloadManager as $MSFT_WebDownloadManager1ref\r\n	{\r\n	ResourceID = \"[ConfigurationRepositoryWeb]AzureAutomationDSC\";\r\n	 SourceInfo = \"C:\\\\OaaS-RegistrationMetaConfig2.ps1::20::9::ConfigurationRepositoryWeb\";\r\n	 RegistrationKey = \"5ci0000000000000000000000000000000000000000000000000000000000000000000000000000Y5H/8wFg==\"; \r\n	 ServerURL = \"https://eus2-agentservice-prod-1.azure-automation.net/accounts/bd8fac9e-0000-0000-0000-0000f474fbf6\";\r\n	};\r\n\r\n	instance of MSFT_WebResourceManager as $MSFT_WebResourceManager1ref\r\n	{\r\n	 SourceInfo = \"C:\\\\OaaS-RegistrationMetaConfig2.ps1::27::9::ResourceRepositoryWeb\";\r\n	 ServerURL = \"https://eus2-agentservice-prod-1.azure-automation.net/accounts/bd8fac9e-0000-0000-0000-0000f474fbf6\";\r\n	 ResourceID = \"[ResourceRepositoryWeb]AzureAutomationDSC\";\r\n	 RegistrationKey = \"5ci0000000000000000000000000000000000000000000000000000000000000000000000000000Y5H/8wFg==\"; \r\n	};\r\n\r\n	instance of MSFT_WebReportManager as $MSFT_WebReportManager1ref\r\n	{\r\n	 SourceInfo = \"C:\\\\OaaS-RegistrationMetaConfig2.ps1::34::9::ReportServerWeb\";\r\n	 ServerURL = \"https://eus2-agentservice-prod-1.azure-automation.net/accounts/bd8fac9e-0000-0000-0000-0000f474fbf6\";\r\n	 ResourceID = \"[ReportServerWeb]AzureAutomationDSC\";\r\n	 RegistrationKey = \"5ci0000000000000000000000000000000000000000000000000000000000000000000000000000Y5H/8wFg==\"; \r\n	};\r\n\r\n	instance of MSFT_DSCMetaConfiguration as $MSFT_DSCMetaConfiguration1ref\r\n	{\r\n	 RefreshMode = \"Pull\";\r\n	 AllowModuleOverwrite = False;\r\n	 ActionAfterReboot = \"ContinueConfiguration\";\r\n	 RefreshFrequencyMins = 30;\r\n	 RebootNodeIfNeeded = False;\r\n	 ConfigurationModeFrequencyMins = 15;\r\n	 ConfigurationMode = \"ApplyAndMonitor\";\r\n\r\n	  ResourceModuleManagers = {\r\n	  $MSFT_WebResourceManager1ref  \r\n	};\r\n	  ReportManagers = {\r\n	  $MSFT_WebReportManager1ref  \r\n	 };\r\n	  ConfigurationDownloadManagers = {\r\n	  $MSFT_WebDownloadManager1ref  \r\n	 };\r\n	};\r\n\r\n	instance of OMI_ConfigurationDocument\r\n	{\r\n	 Version=\"2.0.0\";\r\n	 MinimumCompatibleVersion = \"2.0.0\";\r\n	 CompatibleVersionAdditionalProperties= { \"MSFT_DSCMetaConfiguration:StatusRetentionTimeInDays\" };\r\n	 Author=\"azureautomation\";\r\n	 GenerationDate=\"04/17/2015 11:41:09\";\r\n	 GenerationHost=\"azureautomation-01\";\r\n	 Name=\"RegistrationMetaConfig\";\r\n	};\r\n	"),
	// 	Endpoint: to.Ptr("https://eus2-agentservice-prod-1.azure-automation.net/accounts/bd8fac9e-0000-0000-0000-0000f474fbf6"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount18/agentRegistrationInformation/https://eus2-agentservice-prod-1.azure-automation.net/accounts/bd8fac9e-0000-0000-0000-0000f474fbf6"),
	// 	Keys: &armautomation.AgentRegistrationKeys{
	// 		Primary: to.Ptr("5ci0000000000000000000000000000000000000000000000000000000000000000000000000000Y5H/8wFg=="),
	// 		Secondary: to.Ptr("rVp0000000000000000000000000000000000000000000000000000000000000000000000000000f8cbmrOA=="),
	// 	},
	// }
}
