from azure.identity import DefaultAzureCredential
from azure.mgmt.datalake.store import DataLakeStoreAccountManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-datalake-store
# USAGE
    python accounts_enable_key_vault.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DataLakeStoreAccountManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="34adfa4f-cedf-4dc0-ba29-b6d1a69ab345",
    )

    response = client.accounts.enable_key_vault(
        resource_group_name="contosorg",
        account_name="contosoadla",
    )
    print(response)


# x-ms-original-file: specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/Accounts_EnableKeyVault.json
if __name__ == "__main__":
    main()
