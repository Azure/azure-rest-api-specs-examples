from azure.identity import DefaultAzureCredential

from azure.mgmt.recoveryservicesbackup.activestamp import RecoveryServicesBackupClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-recoveryservicesbackup
# USAGE
    python protection_policies_create_or_update_simple.py

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
                "backupManagementType": "AzureIaasVM",
                "retentionPolicy": {
                    "dailySchedule": {
                        "retentionDuration": {"count": 1, "durationType": "Days"},
                        "retentionTimes": ["2018-01-24T02:00:00Z"],
                    },
                    "retentionPolicyType": "LongTermRetentionPolicy",
                },
                "schedulePolicy": {
                    "schedulePolicyType": "SimpleSchedulePolicy",
                    "scheduleRunFrequency": "Daily",
                    "scheduleRunTimes": ["2018-01-24T02:00:00Z"],
                },
                "timeZone": "Pacific Standard Time",
            }
        },
    )
    print(response)


# x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/AzureIaasVm/ProtectionPolicies_CreateOrUpdate_Simple.json
if __name__ == "__main__":
    main()
