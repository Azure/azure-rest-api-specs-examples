from azure.identity import DefaultAzureCredential
from azure.mgmt.hybridnetwork import HybridNetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-hybridnetwork
# USAGE
    python artifact_store_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = HybridNetworkManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.artifact_stores.begin_create_or_update(
        resource_group_name="rg",
        publisher_name="TestPublisher",
        artifact_store_name="TestArtifactStore",
        parameters={
            "location": "eastus",
            "properties": {
                "managedResourceGroupConfiguration": {"location": "eastus", "name": "testRg"},
                "replicationStrategy": "SingleReplication",
                "storeType": "AzureContainerRegistry",
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/ArtifactStoreCreate.json
if __name__ == "__main__":
    main()
