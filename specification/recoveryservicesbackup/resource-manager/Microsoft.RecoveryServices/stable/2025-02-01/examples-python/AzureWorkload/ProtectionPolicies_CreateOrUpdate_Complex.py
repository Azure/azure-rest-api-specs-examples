from azure.identity import DefaultAzureCredential

from azure.mgmt.recoveryservicesbackup.activestamp import RecoveryServicesBackupClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-recoveryservicesbackup
# USAGE
    python protection_policies_create_or_update_complex.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = RecoveryServicesBackupClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.protection_policies.create_or_update(
        vault_name="NetSDKTestRsVault",
        resource_group_name="SwaggerTestRg",
        policy_name="testPolicy1",
        parameters={
            "properties": {
                "backupManagementType": "AzureWorkload",
                "settings": {"issqlcompression": False, "timeZone": "Pacific Standard Time"},
                "subProtectionPolicy": [
                    {
                        "policyType": "Full",
                        "retentionPolicy": {
                            "monthlySchedule": {
                                "retentionDuration": {"count": 1, "durationType": "Months"},
                                "retentionScheduleFormatType": "Weekly",
                                "retentionScheduleWeekly": {"daysOfTheWeek": ["Sunday"], "weeksOfTheMonth": ["Second"]},
                                "retentionTimes": ["2018-01-24T10:00:00Z"],
                            },
                            "retentionPolicyType": "LongTermRetentionPolicy",
                            "weeklySchedule": {
                                "daysOfTheWeek": ["Sunday", "Tuesday"],
                                "retentionDuration": {"count": 2, "durationType": "Weeks"},
                                "retentionTimes": ["2018-01-24T10:00:00Z"],
                            },
                            "yearlySchedule": {
                                "monthsOfYear": ["January", "June", "December"],
                                "retentionDuration": {"count": 1, "durationType": "Years"},
                                "retentionScheduleFormatType": "Weekly",
                                "retentionScheduleWeekly": {"daysOfTheWeek": ["Sunday"], "weeksOfTheMonth": ["Last"]},
                                "retentionTimes": ["2018-01-24T10:00:00Z"],
                            },
                        },
                        "schedulePolicy": {
                            "schedulePolicyType": "SimpleSchedulePolicy",
                            "scheduleRunDays": ["Sunday", "Tuesday"],
                            "scheduleRunFrequency": "Weekly",
                            "scheduleRunTimes": ["2018-01-24T10:00:00Z"],
                        },
                    },
                    {
                        "policyType": "Differential",
                        "retentionPolicy": {
                            "retentionDuration": {"count": 8, "durationType": "Days"},
                            "retentionPolicyType": "SimpleRetentionPolicy",
                        },
                        "schedulePolicy": {
                            "schedulePolicyType": "SimpleSchedulePolicy",
                            "scheduleRunDays": ["Friday"],
                            "scheduleRunFrequency": "Weekly",
                            "scheduleRunTimes": ["2018-01-24T10:00:00Z"],
                        },
                    },
                    {
                        "policyType": "Log",
                        "retentionPolicy": {
                            "retentionDuration": {"count": 7, "durationType": "Days"},
                            "retentionPolicyType": "SimpleRetentionPolicy",
                        },
                        "schedulePolicy": {"scheduleFrequencyInMins": 60, "schedulePolicyType": "LogSchedulePolicy"},
                    },
                ],
                "workLoadType": "SQLDataBase",
            }
        },
    )
    print(response)


# x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/AzureWorkload/ProtectionPolicies_CreateOrUpdate_Complex.json
if __name__ == "__main__":
    main()
