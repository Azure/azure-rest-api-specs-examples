package armdataprotection_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection"
)

// x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2021-07-01/examples/PolicyCRUD/CreateOrUpdateBackupPolicy.json
func ExampleBackupPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataprotection.NewBackupPoliciesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<vault-name>",
		"<resource-group-name>",
		"<backup-policy-name>",
		armdataprotection.BaseBackupPolicyResource{
			Properties: &armdataprotection.BackupPolicy{
				BaseBackupPolicy: armdataprotection.BaseBackupPolicy{
					DatasourceTypes: []*string{
						to.StringPtr("OssDB")},
					ObjectType: to.StringPtr("<object-type>"),
				},
				PolicyRules: []armdataprotection.BasePolicyRuleClassification{
					&armdataprotection.AzureBackupRule{
						BasePolicyRule: armdataprotection.BasePolicyRule{
							Name:       to.StringPtr("<name>"),
							ObjectType: to.StringPtr("<object-type>"),
						},
						BackupParameters: &armdataprotection.AzureBackupParams{
							BackupParameters: armdataprotection.BackupParameters{
								ObjectType: to.StringPtr("<object-type>"),
							},
							BackupType: to.StringPtr("<backup-type>"),
						},
						DataStore: &armdataprotection.DataStoreInfoBase{
							DataStoreType: armdataprotection.DataStoreTypesVaultStore.ToPtr(),
							ObjectType:    to.StringPtr("<object-type>"),
						},
						Trigger: &armdataprotection.ScheduleBasedTriggerContext{
							TriggerContext: armdataprotection.TriggerContext{
								ObjectType: to.StringPtr("<object-type>"),
							},
							Schedule: &armdataprotection.BackupSchedule{
								RepeatingTimeIntervals: []*string{
									to.StringPtr("R/2019-11-20T08:00:00-08:00/P1W")},
							},
							TaggingCriteria: []*armdataprotection.TaggingCriteria{
								{
									IsDefault: to.BoolPtr(true),
									TagInfo: &armdataprotection.RetentionTag{
										TagName: to.StringPtr("<tag-name>"),
									},
									TaggingPriority: to.Int64Ptr(99),
								},
								{
									Criteria: []armdataprotection.BackupCriteriaClassification{
										&armdataprotection.ScheduleBasedBackupCriteria{
											BackupCriteria: armdataprotection.BackupCriteria{
												ObjectType: to.StringPtr("<object-type>"),
											},
											DaysOfTheWeek: []*armdataprotection.DayOfWeek{
												armdataprotection.DayOfWeekSunday.ToPtr()},
											ScheduleTimes: []*time.Time{
												to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-03-01T13:00:00Z"); return t }())},
										}},
									IsDefault: to.BoolPtr(false),
									TagInfo: &armdataprotection.RetentionTag{
										TagName: to.StringPtr("<tag-name>"),
									},
									TaggingPriority: to.Int64Ptr(20),
								}},
						},
					},
					&armdataprotection.AzureRetentionRule{
						BasePolicyRule: armdataprotection.BasePolicyRule{
							Name:       to.StringPtr("<name>"),
							ObjectType: to.StringPtr("<object-type>"),
						},
						IsDefault: to.BoolPtr(true),
						Lifecycles: []*armdataprotection.SourceLifeCycle{
							{
								DeleteAfter: &armdataprotection.AbsoluteDeleteOption{
									DeleteOption: armdataprotection.DeleteOption{
										Duration:   to.StringPtr("<duration>"),
										ObjectType: to.StringPtr("<object-type>"),
									},
								},
								SourceDataStore: &armdataprotection.DataStoreInfoBase{
									DataStoreType: armdataprotection.DataStoreTypesVaultStore.ToPtr(),
									ObjectType:    to.StringPtr("<object-type>"),
								},
							}},
					},
					&armdataprotection.AzureRetentionRule{
						BasePolicyRule: armdataprotection.BasePolicyRule{
							Name:       to.StringPtr("<name>"),
							ObjectType: to.StringPtr("<object-type>"),
						},
						IsDefault: to.BoolPtr(false),
						Lifecycles: []*armdataprotection.SourceLifeCycle{
							{
								DeleteAfter: &armdataprotection.AbsoluteDeleteOption{
									DeleteOption: armdataprotection.DeleteOption{
										Duration:   to.StringPtr("<duration>"),
										ObjectType: to.StringPtr("<object-type>"),
									},
								},
								SourceDataStore: &armdataprotection.DataStoreInfoBase{
									DataStoreType: armdataprotection.DataStoreTypesVaultStore.ToPtr(),
									ObjectType:    to.StringPtr("<object-type>"),
								},
							}},
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("BaseBackupPolicyResource.ID: %s\n", *res.ID)
}
