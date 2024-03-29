from azure.identity import DefaultAzureCredential
from azure.mgmt.datalake.store import DataLakeStoreAccountManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-datalake-store
# USAGE
    python accounts_create.py

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

    response = client.accounts.begin_create(
        resource_group_name="contosorg",
        account_name="contosoadla",
        parameters={
            "identity": {"type": "SystemAssigned"},
            "location": "eastus2",
            "properties": {
                "defaultGroup": "test_default_group",
                "encryptionConfig": {
                    "keyVaultMetaInfo": {
                        "encryptionKeyName": "test_encryption_key_name",
                        "encryptionKeyVersion": "encryption_key_version",
                        "keyVaultResourceId": "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345",
                    },
                    "type": "UserManaged",
                },
                "encryptionState": "Enabled",
                "firewallAllowAzureIps": "Enabled",
                "firewallRules": [
                    {"name": "test_rule", "properties": {"endIpAddress": "2.2.2.2", "startIpAddress": "1.1.1.1"}}
                ],
                "firewallState": "Enabled",
                "newTier": "Consumption",
                "trustedIdProviderState": "Enabled",
                "trustedIdProviders": [
                    {
                        "name": "test_trusted_id_provider_name",
                        "properties": {"idProvider": "https://sts.windows.net/ea9ec534-a3e3-4e45-ad36-3afc5bb291c1"},
                    }
                ],
            },
            "tags": {"test_key": "test_value"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/Accounts_Create.json
if __name__ == "__main__":
    main()
