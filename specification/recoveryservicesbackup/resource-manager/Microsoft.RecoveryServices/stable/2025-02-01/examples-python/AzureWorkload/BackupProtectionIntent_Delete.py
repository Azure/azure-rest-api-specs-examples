from azure.identity import DefaultAzureCredential

from azure.mgmt.recoveryservicesbackup.activestamp import RecoveryServicesBackupClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-recoveryservicesbackup
# USAGE
    python backup_protection_intent_delete.py

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

    client.protection_intent.delete(
        vault_name="myVault",
        resource_group_name="myRG",
        fabric_name="Azure",
        intent_object_name="249D9B07-D2EF-4202-AA64-65F35418564E",
    )


# x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/AzureWorkload/BackupProtectionIntent_Delete.json
if __name__ == "__main__":
    main()
