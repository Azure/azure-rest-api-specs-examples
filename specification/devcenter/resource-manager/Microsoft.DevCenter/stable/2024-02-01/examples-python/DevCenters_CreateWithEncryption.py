from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.devcenter import DevCenterMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-devcenter
# USAGE
    python dev_centers_create_with_encryption.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DevCenterMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="0ac520ee-14c0-480f-b6c9-0a90c58ffff",
    )

    response = client.dev_centers.begin_create_or_update(
        resource_group_name="rg1",
        dev_center_name="Contoso",
        body={
            "identity": {
                "type": "UserAssigned",
                "userAssignedIdentities": {
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/identityGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testidentity1": {}
                },
            },
            "location": "centralus",
            "properties": {
                "displayName": "ContosoDevCenter",
                "encryption": {
                    "customerManagedKeyEncryption": {
                        "keyEncryptionKeyIdentity": {
                            "identityType": "userAssignedIdentity",
                            "userAssignedIdentityResourceId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/identityGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testidentity1",
                        },
                        "keyEncryptionKeyUrl": "https://contosovault.vault.azure.net/keys/contosokek",
                    }
                },
            },
            "tags": {"CostCode": "12345"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/DevCenters_CreateWithEncryption.json
if __name__ == "__main__":
    main()
