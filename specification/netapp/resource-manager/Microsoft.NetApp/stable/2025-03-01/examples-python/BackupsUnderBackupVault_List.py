from azure.identity import DefaultAzureCredential

from azure.mgmt.netapp import NetAppManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-netapp
# USAGE
    python backups_under_backup_vault_list.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = NetAppManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.backups.list_by_vault(
        resource_group_name="myRG",
        account_name="account1",
        backup_vault_name="backupVault1",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-03-01/examples/BackupsUnderBackupVault_List.json
if __name__ == "__main__":
    main()
