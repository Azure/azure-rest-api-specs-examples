from azure.identity import DefaultAzureCredential

from azure.mgmt.servicefabricmanagedclusters import ServiceFabricManagedClustersManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-servicefabricmanagedclusters
# USAGE
    python node_type_patch_operation_auto_scale_example.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ServiceFabricManagedClustersManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.node_types.begin_update(
        resource_group_name="resRg",
        cluster_name="myCluster",
        node_type_name="BE",
        parameters={"sku": {"capacity": 10, "name": "Standard_S0", "tier": "Standard"}, "tags": {"a": "b"}},
    ).result()
    print(response)


# x-ms-original-file: 2025-03-01-preview/NodeTypePatchOperationAutoScale_example.json
if __name__ == "__main__":
    main()
