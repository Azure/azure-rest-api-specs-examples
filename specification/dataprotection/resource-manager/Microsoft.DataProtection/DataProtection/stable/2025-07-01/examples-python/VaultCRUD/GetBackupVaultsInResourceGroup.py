from azure.identity import DefaultAzureCredential

from azure.mgmt.dataprotection import DataProtectionMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-dataprotection
# USAGE
    python get_backup_vaults_in_resource_group.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DataProtectionMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.backup_vaults.get_in_resource_group(
        resource_group_name="SampleResourceGroup",
    )
    for item in response:
        print(item)


# x-ms-original-file: 2025-07-01/VaultCRUD/GetBackupVaultsInResourceGroup.json
if __name__ == "__main__":
    main()
