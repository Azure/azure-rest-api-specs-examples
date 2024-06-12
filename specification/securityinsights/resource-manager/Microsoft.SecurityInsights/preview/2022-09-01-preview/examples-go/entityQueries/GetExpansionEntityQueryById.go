package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/entityQueries/GetExpansionEntityQueryById.json
func ExampleEntityQueriesClient_Get_getAnExpansionEntityQuery() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEntityQueriesClient().Get(ctx, "myRg", "myWorkspace", "07da3cc8-c8ad-4710-a44e-334cdcb7882b", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.EntityQueriesClientGetResponse{
	// 	                            EntityQueryClassification: &armsecurityinsights.ExpansionEntityQuery{
	// 		Name: to.Ptr("07da3cc8-c8ad-4710-a44e-334cdcb7882b"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/entityQueries"),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/entityQueries/07da3cc8-c8ad-4710-a44e-334cdcb7882b"),
	// 		Kind: to.Ptr(armsecurityinsights.EntityQueryKindExpansion),
	// 		Properties: &armsecurityinsights.ExpansionEntityQueriesProperties{
	// 			DataSources: []*string{
	// 				to.Ptr("SecurityEvent")},
	// 				DisplayName: to.Ptr("Parent processes running on host"),
	// 				InputEntityType: to.Ptr(armsecurityinsights.EntityTypeHost),
	// 				InputFields: []*string{
	// 					to.Ptr("hostName")},
	// 					OutputEntityTypes: []*armsecurityinsights.EntityType{
	// 						to.Ptr(armsecurityinsights.EntityTypeProcess)},
	// 						QueryTemplate: to.Ptr("let GetParentProcessesOnHost = (v_Host_HostName:string){\r\n                            SecurityEvent \r\n                            | where EventID == 4688 \r\n                            | where isnotempty(ParentProcessName)\r\n                            | where NewProcessName !contains ':\\\\Windows\\\\System32\\\\conhost.exe' and ParentProcessName !contains ':\\\\Windows\\\\System32\\\\conhost.exe'\r\n                            and NewProcessName !contains ':\\\\Windows\\\\Microsoft.NET\\\\Framework64\\\\v2.0.50727\\\\csc.exe' and ParentProcessName !contains ':\\\\Windows\\\\Microsoft.NET\\\\Framework64\\\\v2.0.50727\\\\csc.exe'\r\n                            and NewProcessName !contains ':\\\\Windows\\\\Microsoft.NET\\\\Framework64\\\\v2.0.50727\\\\cvtres.exe' and ParentProcessName !contains ':\\\\Windows\\\\Microsoft.NET\\\\Framework64\\\\v2.0.50727\\\\cvtres.exe'\r\n                            and NewProcessName!contains ':\\\\Program Files\\\\Microsoft Monitoring Agent\\\\Agent\\\\MonitoringHost.exe' and ParentProcessName !contains ':\\\\Program Files\\\\Microsoft Monitoring Agent\\\\Agent\\\\MonitoringHost.exe'\r\n                            and ParentProcessName !contains ':\\\\Windows\\\\CCM\\\\CcmExec.exe'\r\n                            | where(ParentProcessName !contains ':\\\\Windows\\\\System32\\\\svchost.exe' and (NewProcessName !contains ':\\\\Windows\\\\System32\\\\wbem\\\\WmiPrvSE.exe' or NewProcessName !contains ':\\\\Windows\\\\SysWOW64\\\\wbem\\\\WmiPrvSE.exe'))\r\n                            | where(ParentProcessName !contains ':\\\\Windows\\\\System32\\\\services.exe' and NewProcessName !contains ':\\\\Windows\\\\servicing\\\\TrustedInstaller.exe')\r\n                            | where toupper(Computer) contains v_Host_HostName or toupper(WorkstationName) contains v_Host_HostName\r\n                            | summarize min(TimeGenerated), max(TimeGenerated) by Account, Computer, ParentProcessName, NewProcessName, CommandLine, ProcessId\r\n                            | project min_TimeGenerated, max_TimeGenerated, Account, Computer, ParentProcessName, NewProcessName, CommandLine, ProcessId\r\n                            | project-rename Process_Host_UnstructuredName=Computer, Process_Account_UnstructuredName=Account, Process_CommandLine=CommandLine, Process_ProcessId=ProcessId, Process_ImageFile_FullPath=NewProcessName, Process_ParentProcess_ImageFile_FullPath=ParentProcessName\r\n                            | top 10 by min_TimeGenerated asc};\r\n                            GetParentProcessesOnHost(toupper('<hostName>'))"),
	// 					},
	// 				},
	// 				                        }
}
