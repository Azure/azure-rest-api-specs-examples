from azure.identity import DefaultAzureCredential
from azure.mgmt.networkcloud import NetworkCloudMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-networkcloud
# USAGE
    python hybrid_aks_clusters_delete.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = NetworkCloudMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="subscriptionId",
    )

    response = client.hybrid_aks_clusters.begin_delete(
        resource_group_name="resourceGroupName",
        hybrid_aks_cluster_name="hybridAksClusterName",
    ).result()
    print(response)


# x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/HybridAksClusters_Delete.json
if __name__ == "__main__":
    main()
