from azure.identity import DefaultAzureCredential

from azure.mgmt.recoveryservicesbackup.activestamp import RecoveryServicesBackupClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-recoveryservicesbackup
# USAGE
    python protection_policies_create_or_update_hourly.py

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
        vault_name="swaggertestvault",
        resource_group_name="SwaggerTestRg",
        policy_name="newPolicy2",
        parameters={
            "properties": {
                "backupManagementType": "AzureStorage",
                "retentionPolicy": {
                    "dailySchedule": {
                        "retentionDuration": {"count": 5, "durationType": "Days"},
                        "retentionTimes": None,
                    },
                    "monthlySchedule": {
                        "retentionDuration": {"count": 60, "durationType": "Months"},
                        "retentionScheduleDaily": None,
                        "retentionScheduleFormatType": "Weekly",
                        "retentionScheduleWeekly": {"daysOfTheWeek": ["Sunday"], "weeksOfTheMonth": ["First"]},
                        "retentionTimes": None,
                    },
                    "retentionPolicyType": "LongTermRetentionPolicy",
                    "weeklySchedule": {
                        "daysOfTheWeek": ["Sunday"],
                        "retentionDuration": {"count": 12, "durationType": "Weeks"},
                        "retentionTimes": None,
                    },
                    "yearlySchedule": {
                        "monthsOfYear": ["January"],
                        "retentionDuration": {"count": 10, "durationType": "Years"},
                        "retentionScheduleDaily": None,
                        "retentionScheduleFormatType": "Weekly",
                        "retentionScheduleWeekly": {"daysOfTheWeek": ["Sunday"], "weeksOfTheMonth": ["First"]},
                        "retentionTimes": None,
                    },
                },
                "schedulePolicy": {
                    "hourlySchedule": {
                        "interval": 4,
                        "scheduleWindowDuration": 12,
                        "scheduleWindowStartTime": "2021-09-29T08:00:00.000Z",
                    },
                    "schedulePolicyType": "SimpleSchedulePolicy",
                    "scheduleRunFrequency": "Hourly",
                },
                "timeZone": "UTC",
                "workLoadType": "AzureFileShare",
            }
        },
    )
    print(response)


# x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/AzureStorage/ProtectionPolicies_CreateOrUpdate_Hourly.json
if __name__ == "__main__":
    main()
