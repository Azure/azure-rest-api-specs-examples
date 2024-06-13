package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/entities/GetQueries.json
func ExampleEntitiesClient_Queries() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEntitiesClient().Queries(ctx, "myRg", "myWorkspace", "e1d3d618-e11f-478b-98e3-bb381539a8e1", armsecurityinsights.EntityItemQueryKindInsight, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GetQueriesResponse = armsecurityinsights.GetQueriesResponse{
	// 	Value: []armsecurityinsights.EntityQueryItemClassification{
	// 		&armsecurityinsights.InsightQueryItem{
	// 			Name: to.Ptr("6db7f5d1-f41e-46c2-b935-230b36a569e6"),
	// 			Type: to.Ptr("Microsoft.SecurityInsights/entities/queries"),
	// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/entities/e1d3d618-e11f-478b-98e3-bb381539a8e1/queries/6db7f5d1-f41e-46c2-b935-230b36a569e6"),
	// 			Kind: to.Ptr(armsecurityinsights.EntityQueryKindInsight),
	// 			Properties: &armsecurityinsights.InsightQueryItemProperties{
	// 				DataTypes: []*armsecurityinsights.EntityQueryItemPropertiesDataTypesItem{
	// 					{
	// 						DataType: to.Ptr("AuditLogs"),
	// 					},
	// 					{
	// 						DataType: to.Ptr("SecurityEvent"),
	// 				}},
	// 				EntitiesFilter: map[string]any{
	// 				},
	// 				InputEntityType: to.Ptr(armsecurityinsights.EntityTypeAccount),
	// 				RequiredInputFieldsSets: [][]*string{
	// 					[]*string{
	// 						to.Ptr("Account_Name"),
	// 						to.Ptr("Account_NTDomain")},
	// 						[]*string{
	// 							to.Ptr("Account_Name"),
	// 							to.Ptr("Account_UPNSuffix")},
	// 							[]*string{
	// 								to.Ptr("Account_AADUserId")},
	// 								[]*string{
	// 									to.Ptr("Account_SID")}},
	// 									Description: to.Ptr("Summary of actions taken on the specified account, grouped by action: password resets and changes, account lockouts (policy or admin), account creation and deletion, account enabled and disabled\n"),
	// 									AdditionalQuery: &armsecurityinsights.InsightQueryItemPropertiesAdditionalQuery{
	// 										Query: to.Ptr("project TimeGenerated, UserPrincipalName, Account_Name, OperationName, Activity, DisableUser, TargetSid, AADUserId, InitiatedBy, AADTenantId, AccountType, Computer, SubjectAccount, SubjectUserSid, EventData"),
	// 										Text: to.Ptr("See all account activity"),
	// 									},
	// 									BaseQuery: to.Ptr("let GetAccountActions = (v_Account_Name:string, v_Account_NTDomain:string, v_Account_UPNSuffix:string, v_Account_AADUserId:string, v_Account_SID:string){\nAuditLogs\n| where OperationName in~ ('Delete user', 'Change user password', 'Reset user password', 'Change password (self-service)',  'Reset password (by admin)', 'Reset password (self-service)', 'Update user')\n| extend UserPrincipalName = tostring(TargetResources[0].userPrincipalName)\n| extend Account_Name = tostring(split(UserPrincipalName, '@')[0])\n| extend Account_UPNSuffix = tostring(split(UserPrincipalName, '@')[1])\n| extend Action = tostring(parse_json(tostring(parse_json(tostring(TargetResources[0].modifiedProperties))[0])))\n| extend ModifiedProperty = parse_json(Action).displayName\n| extend ModifiedValue = parse_json(Action).newValue\n| extend Account_AADUserId = tostring(TargetResources[0].id)\n| extend DisableUser = iif(ModifiedProperty =~ 'AccountEnabled' and ModifiedValue =~ '[false]', 'True', 'False')\n| union isfuzzy=true (\nSecurityEvent\n| where EventID in (4720, 4722, 4723, 4724, 4725, 4726, 4740)\n| extend OperationName = tostring(EventID)\n| where AccountType =~ \"user\" or isempty(AccountType)\n| extend Account_Name = TargetUserName, Account_NTDomain = TargetDomainName, Account_SID = TargetSid\n)\n| where (Account_Name =~ v_Account_Name and (Account_UPNSuffix =~ v_Account_UPNSuffix or Account_NTDomain =~ v_Account_NTDomain)) or Account_AADUserId =~ v_Account_AADUserId or Account_SID =~ v_Account_SID\n};\nGetAccountActions('CTFFUser4', '', 'seccxp.ninja', '', '')\n"),
	// 									ChartQuery: map[string]any{
	// 										"type": "BarChart",
	// 										"dataSets":[]any{
	// 											map[string]any{
	// 												"legendColumnName": "OperationName",
	// 												"query": "summarize Count = count() by bin(TimeGenerated, 1h), OperationName",
	// 												"xColumnName": "TimeGenerated",
	// 												"yColumnName": "Count",
	// 											},
	// 										},
	// 										"title": "Actions by type",
	// 									},
	// 									DefaultTimeRange: &armsecurityinsights.InsightQueryItemPropertiesDefaultTimeRange{
	// 										AfterRange: to.Ptr("12h"),
	// 										BeforeRange: to.Ptr("12h"),
	// 									},
	// 									DisplayName: to.Ptr("Actions on account"),
	// 									TableQuery: &armsecurityinsights.InsightQueryItemPropertiesTableQuery{
	// 										ColumnsDefinitions: []*armsecurityinsights.InsightQueryItemPropertiesTableQueryColumnsDefinitionsItem{
	// 											{
	// 												Header: to.Ptr("Action"),
	// 												OutputType: to.Ptr(armsecurityinsights.OutputTypeString),
	// 												SupportDeepLink: to.Ptr(false),
	// 											},
	// 											{
	// 												Header: to.Ptr("Most Recent"),
	// 												OutputType: to.Ptr(armsecurityinsights.OutputTypeDate),
	// 												SupportDeepLink: to.Ptr(false),
	// 											},
	// 											{
	// 												Header: to.Ptr("Count"),
	// 												OutputType: to.Ptr(armsecurityinsights.OutputTypeNumber),
	// 												SupportDeepLink: to.Ptr(true),
	// 										}},
	// 										QueriesDefinitions: []*armsecurityinsights.InsightQueryItemPropertiesTableQueryQueriesDefinitionsItem{
	// 											{
	// 												Filter: to.Ptr("where OperationName in~ ('Change user password', 'Reset user password', 'Change password (self-service)',  'Reset password (by admin)', 'Reset password (self-service)', '4724', '4723')"),
	// 												LinkColumnsDefinitions: []*armsecurityinsights.InsightQueryItemPropertiesTableQueryQueriesDefinitionsPropertiesItemsItem{
	// 													{
	// 														Query: to.Ptr("{{BaseQuery}} | "),
	// 														ProjectedName: to.Ptr("Count"),
	// 												}},
	// 												Project: to.Ptr("project Title = OperationName, MostRecent, Count"),
	// 												Summarize: to.Ptr("summarize MostRecent = max(TimeGenerated), Count = count() by OperationName"),
	// 											},
	// 											{
	// 												Filter: to.Ptr("where OperationName in~ ('Blocked from self-service password reset', '4740')"),
	// 												LinkColumnsDefinitions: []*armsecurityinsights.InsightQueryItemPropertiesTableQueryQueriesDefinitionsPropertiesItemsItem{
	// 													{
	// 														Query: to.Ptr("{{BaseQuery}} | "),
	// 														ProjectedName: to.Ptr("Count"),
	// 												}},
	// 												Project: to.Ptr("project Title = OperationName, MostRecent, Count"),
	// 												Summarize: to.Ptr("summarize MostRecent = max(TimeGenerated), Count = count() by OperationName"),
	// 											},
	// 											{
	// 												Filter: to.Ptr("where OperationName  == '4725' or (OperationName  =~ 'Update user' and DisableUser =~ 'True')"),
	// 												LinkColumnsDefinitions: []*armsecurityinsights.InsightQueryItemPropertiesTableQueryQueriesDefinitionsPropertiesItemsItem{
	// 													{
	// 														Query: to.Ptr("{{BaseQuery}} | "),
	// 														ProjectedName: to.Ptr("Count"),
	// 												}},
	// 												Project: to.Ptr("project Title = OperationName, MostRecent, Count"),
	// 												Summarize: to.Ptr("summarize MostRecent = max(TimeGenerated), Count = count() by OperationName"),
	// 											},
	// 											{
	// 												Filter: to.Ptr("where OperationName in~ ('Add user', '4720')"),
	// 												LinkColumnsDefinitions: []*armsecurityinsights.InsightQueryItemPropertiesTableQueryQueriesDefinitionsPropertiesItemsItem{
	// 													{
	// 														Query: to.Ptr("{{BaseQuery}} | "),
	// 														ProjectedName: to.Ptr("Count"),
	// 												}},
	// 												Project: to.Ptr("project Title = OperationName, MostRecent, Count"),
	// 												Summarize: to.Ptr("summarize MostRecent = max(TimeGenerated), Count = count() by OperationName"),
	// 											},
	// 											{
	// 												Filter: to.Ptr("where OperationName in~ ('Delete user', '4726')"),
	// 												LinkColumnsDefinitions: []*armsecurityinsights.InsightQueryItemPropertiesTableQueryQueriesDefinitionsPropertiesItemsItem{
	// 													{
	// 														Query: to.Ptr("{{BaseQuery}} | "),
	// 														ProjectedName: to.Ptr("Count"),
	// 												}},
	// 												Project: to.Ptr("project Title = OperationName, MostRecent, Count"),
	// 												Summarize: to.Ptr("summarize MostRecent = max(TimeGenerated), Count = count() by OperationName"),
	// 											},
	// 											{
	// 												Filter: to.Ptr("where OperationName in~ ('4725', 'Blocked from self-service password reset', '4740') or (OperationName  =~ 'Update user' and DisableUser =~ 'True')"),
	// 												LinkColumnsDefinitions: []*armsecurityinsights.InsightQueryItemPropertiesTableQueryQueriesDefinitionsPropertiesItemsItem{
	// 													{
	// 														Query: to.Ptr("{{BaseQuery}} | "),
	// 														ProjectedName: to.Ptr("Count"),
	// 												}},
	// 												Project: to.Ptr("project Title = OperationName, MostRecent, Count"),
	// 												Summarize: to.Ptr("summarize MostRecent = max(TimeGenerated), Count = count() by OperationName"),
	// 											},
	// 											{
	// 												Filter: to.Ptr("where OperationName in~ ('4722', '4767') or (OperationName  =~ 'Update user' and DisableUser =~ 'False')"),
	// 												LinkColumnsDefinitions: []*armsecurityinsights.InsightQueryItemPropertiesTableQueryQueriesDefinitionsPropertiesItemsItem{
	// 													{
	// 														Query: to.Ptr("{{BaseQuery}} | "),
	// 														ProjectedName: to.Ptr("Count"),
	// 												}},
	// 												Project: to.Ptr("project Title = OperationName, MostRecent, Count"),
	// 												Summarize: to.Ptr("summarize MostRecent = max(TimeGenerated), Count = count() by OperationName"),
	// 											},
	// 											{
	// 												Filter: to.Ptr("where OperationName in~ ('Update user','4738')"),
	// 												LinkColumnsDefinitions: []*armsecurityinsights.InsightQueryItemPropertiesTableQueryQueriesDefinitionsPropertiesItemsItem{
	// 													{
	// 														Query: to.Ptr("{{BaseQuery}} | "),
	// 														ProjectedName: to.Ptr("Count"),
	// 												}},
	// 												Project: to.Ptr("project Title = OperationName, MostRecent, Count"),
	// 												Summarize: to.Ptr("summarize MostRecent = max(TimeGenerated), Count = count() by OperationName"),
	// 										}},
	// 									},
	// 								},
	// 							},
	// 							&armsecurityinsights.InsightQueryItem{
	// 								Name: to.Ptr("0a5d7b14-b485-450a-a0ac-4100c860ac32"),
	// 								Type: to.Ptr("Microsoft.SecurityInsights/entities/queries"),
	// 								ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/entities/e1d3d618-e11f-478b-98e3-bb381539a8e1/queries/0a5d7b14-b485-450a-a0ac-4100c860ac32"),
	// 								Kind: to.Ptr(armsecurityinsights.EntityQueryKindInsight),
	// 								Properties: &armsecurityinsights.InsightQueryItemProperties{
	// 									DataTypes: []*armsecurityinsights.EntityQueryItemPropertiesDataTypesItem{
	// 										{
	// 											DataType: to.Ptr("OfficeActivity"),
	// 									}},
	// 									EntitiesFilter: map[string]any{
	// 									},
	// 									InputEntityType: to.Ptr(armsecurityinsights.EntityTypeAccount),
	// 									RequiredInputFieldsSets: [][]*string{
	// 										[]*string{
	// 											to.Ptr("Account_Name"),
	// 											to.Ptr("Account_UPNSuffix")}},
	// 											Description: to.Ptr("Highlight office operations of the user with anomalously high count compared to those observed in the preceding 14 days."),
	// 											AdditionalQuery: &armsecurityinsights.InsightQueryItemPropertiesAdditionalQuery{
	// 												Query: to.Ptr("make-series count() default=0 on TimeGenerated from (StartTime - BeforeRange) to EndTime step 1d by Operation \n| extend (anomalies,anomalyScore, expectedCount)=series_decompose_anomalies(count_,AScoreThresh,7,'linefit',numDays, 'ctukey') \n| extend count1=count_, TimeGenerated1=TimeGenerated, anomalyScore1=anomalyScore\n| mv-apply count1 to typeof(long), TimeGenerated1 to typeof(datetime), anomalyScore1 to typeof(double), anomalies to typeof(long) on (summarize totAnomalies=sumif(abs(anomalies), TimeGenerated1 < StartTime), baseStd=stdevif(count1, TimeGenerated1 < StartTime), baseAvg=avgif(count1, TimeGenerated1 < StartTime), maxCountPost=maxif(count1,TimeGenerated1 >= StartTime), maxAnomalyScorePost = maxif(anomalyScore1, TimeGenerated1 >= StartTime)) \n| extend count1=count_\n| mv-apply  count1 to typeof(long), anomalyScore to typeof(double), expectedCount to typeof(double) on ( summarize (dummy, postExpectedCount, postActualCount)=arg_min(abs(anomalyScore - maxAnomalyScorePost), expectedCount, count1) ) \n| where totAnomalies < maxAnomalies\n| extend postAnomalyScore=iff(baseStd == 0 and maxCountPost > tolong(count_[0]),1000.0,maxAnomalyScorePost), postExpectedCount=iff(postExpectedCount < 0,0.0,postExpectedCount) \n| where maxAnomalyScorePost > AScoreThresh | order by maxAnomalyScorePost desc \n| project Operation, expectedCount=round(postExpectedCount,2), actualCount=postActualCount, anomalyScore=round(postAnomalyScore,2)\n"),
	// 												Text: to.Ptr("Query all anomalous operations"),
	// 											},
	// 											BaseQuery: to.Ptr("let AScoreThresh = 3; \nlet maxAnomalies = 3;\nlet BeforeRange = 12d; \nlet EndTime = todatetime('{{EndTimeUTC}}'); \nlet StartTime = todatetime('{{StartTimeUTC}}');\nlet numDays = tolong((EndTime-StartTime)/1d); \nlet userData = (v_Account_Name:string, v_Account_UPNSuffix:string) { \n  OfficeActivity \n  | extend splitUserId=split(UserId, '@')\n  | extend Account_Name = tostring(splitUserId[0]), Account_UPNSuffix = tostring(splitUserId[1])\n  | where Account_Name =~ v_Account_Name and Account_UPNSuffix =~ v_Account_UPNSuffix }; \nuserData('CTFFUser4', 'seccxp.ninja')\n"),
	// 											ChartQuery: map[string]any{
	// 												"type": "LineChart",
	// 												"dataSets":[]any{
	// 													map[string]any{
	// 														"legendColumnName": "Operation",
	// 														"query": "make-series count() default=0 on TimeGenerated from (StartTime - BeforeRange) to EndTime step 1d by Operation \n| extend (anomalies,anomalyScore, expectedCount)=series_decompose_anomalies(count_,AScoreThresh,7,'linefit',numDays, 'ctukey') \n| extend count1=count_, TimeGenerated1=TimeGenerated, anomalyScore1=anomalyScore\n| mv-apply count1 to typeof(long), TimeGenerated1 to typeof(datetime), anomalyScore1 to typeof(double), anomalies to typeof(long) on (summarize totAnomalies=sumif(abs(anomalies), TimeGenerated1 < StartTime), baseStd=stdevif(count1, TimeGenerated1 < StartTime), baseAvg=avgif(count1, TimeGenerated1 < StartTime), maxCountPost=maxif(count1,TimeGenerated1 >= StartTime), maxAnomalyScorePost=maxif(anomalyScore1, TimeGenerated1 >= StartTime)) \n| extend count1=count_ \n| mv-apply count1 to typeof(long), anomalyScore to typeof(double), expectedCount to typeof(double) on ( summarize (dummy, postExpectedCount, postActualCount)=arg_min(abs(anomalyScore-maxAnomalyScorePost), expectedCount, count1) ) \n| where totAnomalies < maxAnomalies \n| extend postAnomalyScore=iff(baseStd == 0 and maxCountPost > tolong(count_[0]),1000.0,maxAnomalyScorePost), postExpectedCount=iff(postExpectedCount < 0,0.0,round(postExpectedCount,2)) \n| where maxAnomalyScorePost > AScoreThresh \n| order by maxAnomalyScorePost desc \n| take 1 \n| project Operation, TimeGenerated, count_\n| mvexpand TimeGenerated, count_ | project todatetime(TimeGenerated), toint(count_), Operation\n",
	// 														"xColumnName": "TimeGenerated",
	// 														"yColumnName": "count_",
	// 													},
	// 												},
	// 												"title": "Anomalous operation timeline",
	// 											},
	// 											DefaultTimeRange: &armsecurityinsights.InsightQueryItemPropertiesDefaultTimeRange{
	// 												AfterRange: to.Ptr("0d"),
	// 												BeforeRange: to.Ptr("1d"),
	// 											},
	// 											DisplayName: to.Ptr("Anomalously high office operation count"),
	// 											ReferenceTimeRange: &armsecurityinsights.InsightQueryItemPropertiesReferenceTimeRange{
	// 												BeforeRange: to.Ptr("12d"),
	// 											},
	// 											TableQuery: &armsecurityinsights.InsightQueryItemPropertiesTableQuery{
	// 												ColumnsDefinitions: []*armsecurityinsights.InsightQueryItemPropertiesTableQueryColumnsDefinitionsItem{
	// 													{
	// 														Header: to.Ptr("Operation"),
	// 														OutputType: to.Ptr(armsecurityinsights.OutputTypeString),
	// 														SupportDeepLink: to.Ptr(true),
	// 													},
	// 													{
	// 														Header: to.Ptr("Expected Count"),
	// 														OutputType: to.Ptr(armsecurityinsights.OutputTypeNumber),
	// 														SupportDeepLink: to.Ptr(false),
	// 													},
	// 													{
	// 														Header: to.Ptr("Actual Count"),
	// 														OutputType: to.Ptr(armsecurityinsights.OutputTypeNumber),
	// 														SupportDeepLink: to.Ptr(false),
	// 												}},
	// 												QueriesDefinitions: []*armsecurityinsights.InsightQueryItemPropertiesTableQueryQueriesDefinitionsItem{
	// 													{
	// 														Filter: to.Ptr("make-series count() default=0 on TimeGenerated from (StartTime - BeforeRange) to EndTime step 1d by Operation \n| extend (anomalies,anomalyScore, expectedCount)=series_decompose_anomalies(count_,AScoreThresh,7,'linefit',numDays, 'ctukey') \n| extend count1=count_, TimeGenerated1=TimeGenerated, anomalyScore1=anomalyScore\n| mv-apply count1 to typeof(long), TimeGenerated1 to typeof(datetime), anomalyScore1 to typeof(double), anomalies to typeof(long) on (summarize totAnomalies=sumif(abs(anomalies), TimeGenerated1 < StartTime), baseStd=stdevif(count1, TimeGenerated1 < StartTime), baseAvg=avgif(count1, TimeGenerated1 < StartTime), maxCountPost=maxif(count1,TimeGenerated1 >= StartTime), maxAnomalyScorePost=maxif(anomalyScore1, TimeGenerated1 >= StartTime)) \n| extend count1=count_ \n| mv-apply count1 to typeof(long), anomalyScore to typeof(double), expectedCount to typeof(double) on ( summarize (dummy, postExpectedCount, postActualCount)=arg_min(abs(anomalyScore-maxAnomalyScorePost), expectedCount, count1) ) \n| where totAnomalies < maxAnomalies \n| extend postAnomalyScore=iff(baseStd == 0 and maxCountPost > tolong(count_[0]),1000.0,maxAnomalyScorePost), postExpectedCount=iff(postExpectedCount < 0,0.0,postExpectedCount) \n| where maxAnomalyScorePost > AScoreThresh \n| order by maxAnomalyScorePost desc\n"),
	// 														LinkColumnsDefinitions: []*armsecurityinsights.InsightQueryItemPropertiesTableQueryQueriesDefinitionsPropertiesItemsItem{
	// 															{
	// 																Query: to.Ptr("{{BaseQuery}} \n| where TimeGenerated between (StartTime .. EndTime) \n| where Operation == ''\n"),
	// 																ProjectedName: to.Ptr("Operation"),
	// 														}},
	// 														Project: to.Ptr("project Operation, expectedCount=round(postExpectedCount,2), actualCount=postActualCount, anomalyScore=round(postAnomalyScore,2)"),
	// 														Summarize: to.Ptr("take 1"),
	// 												}},
	// 											},
	// 										},
	// 									},
	// 									&armsecurityinsights.InsightQueryItem{
	// 										Name: to.Ptr("e6cf68e6-1eca-4fbb-9fad-6280f2a9476e"),
	// 										Type: to.Ptr("Microsoft.SecurityInsights/entities/queries"),
	// 										ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/entities/e1d3d618-e11f-478b-98e3-bb381539a8e1/queries/e6cf68e6-1eca-4fbb-9fad-6280f2a9476e"),
	// 										Kind: to.Ptr(armsecurityinsights.EntityQueryKindInsight),
	// 										Properties: &armsecurityinsights.InsightQueryItemProperties{
	// 											DataTypes: []*armsecurityinsights.EntityQueryItemPropertiesDataTypesItem{
	// 												{
	// 													DataType: to.Ptr("OfficeActivity"),
	// 											}},
	// 											EntitiesFilter: map[string]any{
	// 											},
	// 											InputEntityType: to.Ptr(armsecurityinsights.EntityTypeAccount),
	// 											RequiredInputFieldsSets: [][]*string{
	// 												[]*string{
	// 													to.Ptr("Account_Name"),
	// 													to.Ptr("Account_UPNSuffix")},
	// 													[]*string{
	// 														to.Ptr("Account_AADUserId")}},
	// 														Description: to.Ptr("Provides the count and distinct resource accesses by a given user account\n"),
	// 														AdditionalQuery: &armsecurityinsights.InsightQueryItemPropertiesAdditionalQuery{
	// 															Query: to.Ptr("where Operation in~ (Operations)"),
	// 															Text: to.Ptr("See all resource activity"),
	// 														},
	// 														BaseQuery: to.Ptr("let Operations = dynamic([\"FileDownloaded\", \"FileUploaded\"]);\nlet UserOperationToSharePoint =  (v_Account_Name:string, v_Account_UPNSuffix:string) {\nOfficeActivity\n// Select sharepoint activity that is relevant\n| where RecordType in~ ('SharePointFileOperation')\n| where Operation in~ (Operations)\n| extend Account_Name = tostring(split(UserId, '@')[0])\n| extend Account_UPNSuffix = tostring(split(UserId, '@')[1])\n| where Account_Name =~ v_Account_Name and Account_UPNSuffix =~ v_Account_UPNSuffix\n| project TimeGenerated, Account_Name, Account_UPNSuffix, UserId, OfficeId, RecordType, Operation, OrganizationId, UserType, UserKey, OfficeWorkload, OfficeObjectId, ClientIP, ItemType, UserAgent, Site_Url, SourceRelativeUrl, SourceFileName, SourceFileExtension , Start_Time , ElevationTime , TenantId, SourceSystem , Type\n};\nUserOperationToSharePoint ('CTFFUser4','seccxp.ninja')\n"),
	// 														ChartQuery: map[string]any{
	// 															"type": "LineChart",
	// 															"dataSets":[]any{
	// 																map[string]any{
	// 																	"legendColumnName": "Legend",
	// 																	"query": "summarize DistinctResources = dcountif(Operation, Operation =~ 'FileUploaded'), TotalResources = countif(Operation =~ 'FileUploaded') by bin(TimeGenerated, 1h) | extend Legend = 'File Uploads'",
	// 																	"xColumnName": "TimeGenerated",
	// 																	"yColumnName": "TotalResources",
	// 																},
	// 																map[string]any{
	// 																	"legendColumnName": "Legend",
	// 																	"query": "summarize DistinctResources = dcountif(Operation, Operation =~ 'FileDownloaded'), TotalResources = countif(Operation =~ 'FileDownloaded') by bin(TimeGenerated, 1h) | extend Legend = 'File Downloads'",
	// 																	"xColumnName": "TimeGenerated",
	// 																	"yColumnName": "TotalResources",
	// 																},
	// 															},
	// 															"title": "Resource access over time",
	// 														},
	// 														DefaultTimeRange: &armsecurityinsights.InsightQueryItemPropertiesDefaultTimeRange{
	// 															AfterRange: to.Ptr("12h"),
	// 															BeforeRange: to.Ptr("12h"),
	// 														},
	// 														DisplayName: to.Ptr("Resource access"),
	// 														TableQuery: &armsecurityinsights.InsightQueryItemPropertiesTableQuery{
	// 															ColumnsDefinitions: []*armsecurityinsights.InsightQueryItemPropertiesTableQueryColumnsDefinitionsItem{
	// 																{
	// 																	Header: to.Ptr("Resource Type"),
	// 																	OutputType: to.Ptr(armsecurityinsights.OutputTypeString),
	// 																	SupportDeepLink: to.Ptr(false),
	// 																},
	// 																{
	// 																	Header: to.Ptr("Distinct Resources"),
	// 																	OutputType: to.Ptr(armsecurityinsights.OutputTypeNumber),
	// 																	SupportDeepLink: to.Ptr(true),
	// 																},
	// 																{
	// 																	Header: to.Ptr("Total Resources"),
	// 																	OutputType: to.Ptr(armsecurityinsights.OutputTypeNumber),
	// 																	SupportDeepLink: to.Ptr(true),
	// 																},
	// 																{
	// 																	Header: to.Ptr("IPAddress(es)"),
	// 																	OutputType: to.Ptr(armsecurityinsights.OutputTypeString),
	// 																	SupportDeepLink: to.Ptr(false),
	// 															}},
	// 															QueriesDefinitions: []*armsecurityinsights.InsightQueryItemPropertiesTableQueryQueriesDefinitionsItem{
	// 																{
	// 																	Filter: to.Ptr("where Operation =~ 'FileUploaded'"),
	// 																	LinkColumnsDefinitions: []*armsecurityinsights.InsightQueryItemPropertiesTableQueryQueriesDefinitionsPropertiesItemsItem{
	// 																		{
	// 																			Query: to.Ptr("{{BaseQuery}} | "),
	// 																			ProjectedName: to.Ptr("DistinctResources"),
	// 																		},
	// 																		{
	// 																			Query: to.Ptr("{{BaseQuery}} | "),
	// 																			ProjectedName: to.Ptr("TotalResources"),
	// 																	}},
	// 																	Project: to.Ptr("project Title = Operation, DistinctResources, TotalResources, IPAddresses = case(array_length(IPAddresses) == 1, tostring(IPAddresses[0]), array_length(IPAddresses) > 1, 'Many', 'None')"),
	// 																	Summarize: to.Ptr("summarize DistinctResources = dcount(SourceFileName), TotalResources = count(SourceFileName), IPAddresses = make_set(ClientIP) by Operation"),
	// 																},
	// 																{
	// 																	Filter: to.Ptr("where Operation =~ 'FileDownloaded'"),
	// 																	LinkColumnsDefinitions: []*armsecurityinsights.InsightQueryItemPropertiesTableQueryQueriesDefinitionsPropertiesItemsItem{
	// 																		{
	// 																			Query: to.Ptr("{{BaseQuery}} | "),
	// 																			ProjectedName: to.Ptr("DistinctResources"),
	// 																		},
	// 																		{
	// 																			Query: to.Ptr("{{BaseQuery}} | "),
	// 																			ProjectedName: to.Ptr("TotalResources"),
	// 																	}},
	// 																	Project: to.Ptr("project Title = Operation, DistinctResources, TotalResources, IPAddresses = case(array_length(IPAddresses) == 1, tostring(IPAddresses[0]), array_length(IPAddresses) > 1, 'Many', 'None')"),
	// 																	Summarize: to.Ptr("summarize DistinctResources = dcount(SourceFileName), TotalResources = count(SourceFileName), IPAddresses = make_set(ClientIP) by Operation"),
	// 															}},
	// 														},
	// 													},
	// 												},
	// 												&armsecurityinsights.InsightQueryItem{
	// 													Name: to.Ptr("cae8d0aa-aa45-4d53-8d88-17dd64ffd4e4"),
	// 													Type: to.Ptr("Microsoft.SecurityInsights/entities/queries"),
	// 													ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/entities/e1d3d618-e11f-478b-98e3-bb381539a8e1/queries/cae8d0aa-aa45-4d53-8d88-17dd64ffd4e4"),
	// 													Kind: to.Ptr(armsecurityinsights.EntityQueryKindInsight),
	// 													Properties: &armsecurityinsights.InsightQueryItemProperties{
	// 														DataTypes: []*armsecurityinsights.EntityQueryItemPropertiesDataTypesItem{
	// 															{
	// 																DataType: to.Ptr("SigninLogs"),
	// 														}},
	// 														EntitiesFilter: map[string]any{
	// 														},
	// 														InputEntityType: to.Ptr(armsecurityinsights.EntityTypeAccount),
	// 														RequiredInputFieldsSets: [][]*string{
	// 															[]*string{
	// 																to.Ptr("Account_Name"),
	// 																to.Ptr("Account_UPNSuffix")},
	// 																[]*string{
	// 																	to.Ptr("Account_AADUserId")}},
	// 																	Description: to.Ptr("Highlight Azure sign-in results by the user principal with anomalously high count compared to those observed in the preceding 14 days."),
	// 																	AdditionalQuery: &armsecurityinsights.InsightQueryItemPropertiesAdditionalQuery{
	// 																		Query: to.Ptr("make-series count() default=0 on TimeGenerated from (StartTime - BeforeRange) to EndTime step 1d by ResultDescription \n| extend (anomalies,anomalyScore, expectedCount)=series_decompose_anomalies(count_,AScoreThresh,7,'linefit',numDays, 'ctukey') \n| extend count1=count_, TimeGenerated1=TimeGenerated, anomalyScore1=anomalyScore\n| mv-apply count1 to typeof(long), TimeGenerated1 to typeof(datetime), anomalyScore1 to typeof(double), anomalies to typeof(long) on (summarize totAnomalies=sumif(abs(anomalies), TimeGenerated1 < StartTime), baseStd=stdevif(count1, TimeGenerated1 < StartTime), baseAvg=avgif(count1, TimeGenerated1 < StartTime), maxCountPost=maxif(count1,TimeGenerated1 >= StartTime), maxAnomalyScorePost = maxif(anomalyScore1, TimeGenerated1 >= StartTime)) \n| extend count1=count_\n| mv-apply  count1 to typeof(long), anomalyScore to typeof(double), expectedCount to typeof(double) on ( summarize (dummy, postExpectedCount, postActualCount)=arg_min(abs(anomalyScore - maxAnomalyScorePost), expectedCount, count1) ) \n| where totAnomalies < maxAnomalies\n| extend postAnomalyScore=iff(baseStd == 0 and maxCountPost > tolong(count_[0]),1000.0,maxAnomalyScorePost), postExpectedCount=iff(postExpectedCount < 0,0.0,postExpectedCount) \n| where maxAnomalyScorePost > AScoreThresh \n| order by maxAnomalyScorePost desc \n| project ResultDescription, expectedCount=round(postExpectedCount,2), actualCount=postActualCount, anomalyScore=round(postAnomalyScore,2)\n"),
	// 																		Text: to.Ptr("Query all anomalous sign-in results"),
	// 																	},
	// 																	BaseQuery: to.Ptr("let AScoreThresh=3; \nlet maxAnomalies=3; \nlet BeforeRange = 12d; \nlet EndTime=todatetime('{{EndTimeUTC}}');\nlet StartTime = todatetime('{{StartTimeUTC}}'); \nlet numDays = tolong((EndTime-StartTime)/1d); \nlet userData = (v_Account_Name:string, v_Account_UPNSuffix:string, v_Account_AADUserId:string) { \n   SigninLogs \n   | where TimeGenerated between ((StartTime-BeforeRange) .. EndTime)\n   | extend splitUserId=split(UserPrincipalName, '@')\n   | extend Account_Name = tostring(splitUserId[0]), Account_UPNSuffix = tostring(splitUserId[1])\n   | where (Account_Name =~ v_Account_Name and Account_UPNSuffix =~ v_Account_UPNSuffix) or UserId =~ v_Account_AADUserId };\nuserData('CTFFUser4', 'seccxp.ninja', '')\n"),
	// 																	ChartQuery: map[string]any{
	// 																		"type": "LineChart",
	// 																		"dataSets":[]any{
	// 																			map[string]any{
	// 																				"legendColumnName": "ResultDescription",
	// 																				"query": "make-series count() default=0 on TimeGenerated from (StartTime - BeforeRange) to EndTime step 1d by ResultDescription \n| extend (anomalies,anomalyScore, expectedCount)=series_decompose_anomalies(count_,AScoreThresh,7,'linefit',numDays, 'ctukey') \n| extend count1=count_, TimeGenerated1=TimeGenerated, anomalyScore1=anomalyScore\n| mv-apply count1 to typeof(long), TimeGenerated1 to typeof(datetime), anomalyScore1 to typeof(double), anomalies to typeof(long) on (summarize totAnomalies=sumif(abs(anomalies), TimeGenerated1 < StartTime), baseStd=stdevif(count1, TimeGenerated1 < StartTime), baseAvg=avgif(count1, TimeGenerated1 < StartTime), maxCountPost=maxif(count1,TimeGenerated1 >= StartTime), maxAnomalyScorePost = maxif(anomalyScore1, TimeGenerated1 >= StartTime)) \n| extend count1=count_ \n| mv-apply  count1 to typeof(long), anomalyScore to typeof(double), expectedCount to typeof(double) on ( summarize (dummy, postExpectedCount, postActualCount)=arg_min(abs(anomalyScore - maxAnomalyScorePost), expectedCount, count1) ) \n| where totAnomalies < maxAnomalies \n| extend postAnomalyScore=iff(baseStd == 0 and maxCountPost > tolong(count_[0]),1000.0,maxAnomalyScorePost), postExpectedCount=iff(postExpectedCount < 0,0.0,round(postExpectedCount,2)) \n| where maxAnomalyScorePost > AScoreThresh \n| order by maxAnomalyScorePost desc \n| take 1 \n| project ResultDescription, TimeGenerated, count_ \n| mvexpand TimeGenerated, count_ \n| project todatetime(TimeGenerated), toint(count_), ResultDescription \n",
	// 																				"xColumnName": "TimeGenerated",
	// 																				"yColumnName": "count_",
	// 																			},
	// 																		},
	// 																		"title": "Anomalous sign-in result timeline",
	// 																	},
	// 																	DefaultTimeRange: &armsecurityinsights.InsightQueryItemPropertiesDefaultTimeRange{
	// 																		AfterRange: to.Ptr("0d"),
	// 																		BeforeRange: to.Ptr("1d"),
	// 																	},
	// 																	DisplayName: to.Ptr("Anomalously high Azure sign-in result count"),
	// 																	ReferenceTimeRange: &armsecurityinsights.InsightQueryItemPropertiesReferenceTimeRange{
	// 																		BeforeRange: to.Ptr("12d"),
	// 																	},
	// 																	TableQuery: &armsecurityinsights.InsightQueryItemPropertiesTableQuery{
	// 																		ColumnsDefinitions: []*armsecurityinsights.InsightQueryItemPropertiesTableQueryColumnsDefinitionsItem{
	// 																			{
	// 																				Header: to.Ptr("Result Description"),
	// 																				OutputType: to.Ptr(armsecurityinsights.OutputTypeString),
	// 																				SupportDeepLink: to.Ptr(true),
	// 																			},
	// 																			{
	// 																				Header: to.Ptr("Expected Count"),
	// 																				OutputType: to.Ptr(armsecurityinsights.OutputTypeNumber),
	// 																				SupportDeepLink: to.Ptr(false),
	// 																			},
	// 																			{
	// 																				Header: to.Ptr("Actual Count"),
	// 																				OutputType: to.Ptr(armsecurityinsights.OutputTypeNumber),
	// 																				SupportDeepLink: to.Ptr(false),
	// 																		}},
	// 																		QueriesDefinitions: []*armsecurityinsights.InsightQueryItemPropertiesTableQueryQueriesDefinitionsItem{
	// 																			{
	// 																				Filter: to.Ptr("make-series count() default=0 on TimeGenerated from (StartTime - BeforeRange) to EndTime step 1d by ResultDescription \n| extend (anomalies,anomalyScore, expectedCount)=series_decompose_anomalies(count_,AScoreThresh,7,'linefit',numDays, 'ctukey') \n| extend count1=count_, TimeGenerated1=TimeGenerated, anomalyScore1=anomalyScore\n| mv-apply count1 to typeof(long), TimeGenerated1 to typeof(datetime), anomalyScore1 to typeof(double), anomalies to typeof(long) on (summarize totAnomalies=sumif(abs(anomalies), TimeGenerated1 < StartTime), baseStd=stdevif(count1, TimeGenerated1 < StartTime), baseAvg=avgif(count1, TimeGenerated1 < StartTime), maxCountPost=maxif(count1,TimeGenerated1 >= StartTime), maxAnomalyScorePost = maxif(anomalyScore1, TimeGenerated1 >= StartTime)) \n| extend count1=count_ \n| mv-apply  count1 to typeof(long), anomalyScore to typeof(double), expectedCount to typeof(double) on ( summarize (dummy, postExpectedCount, postActualCount)=arg_min(abs(anomalyScore - maxAnomalyScorePost), expectedCount, count1) ) \n| where totAnomalies < maxAnomalies \n| extend postAnomalyScore=iff(baseStd == 0 and maxCountPost > tolong(count_[0]),1000.0,maxAnomalyScorePost), postExpectedCount=iff(postExpectedCount < 0,0.0,postExpectedCount) \n| where maxAnomalyScorePost > AScoreThresh \n| order by maxAnomalyScorePost desc\n"),
	// 																				LinkColumnsDefinitions: []*armsecurityinsights.InsightQueryItemPropertiesTableQueryQueriesDefinitionsPropertiesItemsItem{
	// 																					{
	// 																						Query: to.Ptr("{{BaseQuery}} \n| where TimeGenerated between (StartTime .. EndTime) \n| where ResultDescription == ''\n"),
	// 																						ProjectedName: to.Ptr("ResultDescription"),
	// 																				}},
	// 																				Project: to.Ptr("project ResultDescription, expectedCount=round(postExpectedCount,2), actualCount=postActualCount, anomalyScore=round(postAnomalyScore,2)"),
	// 																				Summarize: to.Ptr("take 1"),
	// 																		}},
	// 																	},
	// 																},
	// 														}},
	// 													}
}
