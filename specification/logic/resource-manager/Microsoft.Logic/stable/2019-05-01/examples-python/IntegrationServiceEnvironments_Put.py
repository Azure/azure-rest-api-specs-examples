from azure.identity import DefaultAzureCredential
from azure.mgmt.logic import LogicManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-logic
# USAGE
    python create_or_update_an_integration_service_environment.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = LogicManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="f34b22a3-2202-4fb1-b040-1332bd928c84",
    )

    response = client.integration_service_environments.begin_create_or_update(
        resource_group="testResourceGroup",
        integration_service_environment_name="testIntegrationServiceEnvironment",
        integration_service_environment={
            "location": "brazilsouth",
            "properties": {
                "encryptionConfiguration": {
                    "encryptionKeyReference": {
                        "keyName": "testKeyName",
                        "keyVault": {
                            "id": "/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.KeyVault/vaults/testKeyVault"
                        },
                        "keyVersion": "13b261d30b984753869902d7f47f4d55",
                    }
                },
                "networkConfiguration": {
                    "accessEndpoint": {"type": "Internal"},
                    "subnets": [
                        {
                            "id": "/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Network/virtualNetworks/testVNET/subnets/s1"
                        },
                        {
                            "id": "/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Network/virtualNetworks/testVNET/subnets/s2"
                        },
                        {
                            "id": "/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Network/virtualNetworks/testVNET/subnets/s3"
                        },
                        {
                            "id": "/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Network/virtualNetworks/testVNET/subnets/s4"
                        },
                    ],
                },
            },
            "sku": {"capacity": 2, "name": "Premium"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationServiceEnvironments_Put.json
if __name__ == "__main__":
    main()
