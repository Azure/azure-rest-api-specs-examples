from azure.identity import DefaultAzureCredential
from azure.mgmt.machinelearningservices import MachineLearningServicesMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-machinelearningservices
# USAGE
    python create_or_update_system_created.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = MachineLearningServicesMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-1111-2222-3333-444444444444",
    )

    response = client.registries.begin_create_or_update(
        resource_group_name="test-rg",
        registry_name="string",
        body={
            "identity": {"type": "None", "userAssignedIdentities": {"string": {}}},
            "kind": "string",
            "location": "string",
            "properties": {
                "discoveryUrl": "string",
                "intellectualPropertyPublisher": "string",
                "managedResourceGroup": {"resourceId": "string"},
                "mlFlowRegistryUri": "string",
                "publicNetworkAccess": "string",
                "regionDetails": [
                    {
                        "acrDetails": [
                            {
                                "systemCreatedAcrAccount": {
                                    "acrAccountName": "string",
                                    "acrAccountSku": "string",
                                    "armResourceId": {"resourceId": "string"},
                                }
                            }
                        ],
                        "location": "string",
                        "storageAccountDetails": [
                            {
                                "systemCreatedStorageAccount": {
                                    "allowBlobPublicAccess": False,
                                    "armResourceId": {"resourceId": "string"},
                                    "storageAccountHnsEnabled": False,
                                    "storageAccountName": "string",
                                    "storageAccountType": "string",
                                }
                            }
                        ],
                    }
                ],
                "registryPrivateEndpointConnections": [
                    {
                        "id": "string",
                        "location": "string",
                        "properties": {
                            "groupIds": ["string"],
                            "privateEndpoint": {"subnetArmId": "string"},
                            "provisioningState": "string",
                            "registryPrivateLinkServiceConnectionState": {
                                "actionsRequired": "string",
                                "description": "string",
                                "status": "Approved",
                            },
                        },
                    }
                ],
            },
            "sku": {"capacity": 1, "family": "string", "name": "string", "size": "string", "tier": "Free"},
            "tags": {},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2023-04-01/examples/Registries/createOrUpdate-SystemCreated.json
if __name__ == "__main__":
    main()
