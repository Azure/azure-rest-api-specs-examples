from azure.identity import DefaultAzureCredential

from azure.mgmt.storage import StorageManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-storage
# USAGE
    python storage_account_create_user_assigned_identity_with_federated_identity_client_id.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = StorageManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="{subscription-id}",
    )

    response = client.storage_accounts.begin_create(
        resource_group_name="res131918",
        account_name="sto131918",
        parameters={
            "identity": {
                "type": "UserAssigned",
                "userAssignedIdentities": {
                    "/subscriptions/{subscription-id}/resourceGroups/res9101/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{managed-identity-name}": {}
                },
            },
            "kind": "Storage",
            "location": "eastus",
            "properties": {
                "encryption": {
                    "identity": {
                        "federatedIdentityClientId": "f83c6b1b-4d34-47e4-bb34-9d83df58b540",
                        "userAssignedIdentity": "/subscriptions/{subscription-id}/resourceGroups/res9101/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{managed-identity-name}",
                    },
                    "keySource": "Microsoft.Keyvault",
                    "keyvaultproperties": {
                        "keyname": "wrappingKey",
                        "keyvaulturi": "https://myvault8569.vault.azure.net",
                        "keyversion": "",
                    },
                    "services": {
                        "blob": {"enabled": True, "keyType": "Account"},
                        "file": {"enabled": True, "keyType": "Account"},
                    },
                }
            },
            "sku": {"name": "Standard_LRS"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/StorageAccountCreateUserAssignedIdentityWithFederatedIdentityClientId.json
if __name__ == "__main__":
    main()
