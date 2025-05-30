from azure.identity import DefaultAzureCredential

from azure.mgmt.servicefabricmanagedclusters import ServiceFabricManagedClustersManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-servicefabricmanagedclusters
# USAGE
    python node_type_delete_operation_example.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ServiceFabricManagedClustersManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    client.node_types.begin_delete(
        resource_group_name="resRg",
        cluster_name="myCluster",
        node_type_name="BE",
    ).result()


# x-ms-original-file: specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2024-06-01-preview/examples/NodeTypeDeleteOperation_example.json
if __name__ == "__main__":
    main()
