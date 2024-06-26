from azure.identity import DefaultAzureCredential
from azure.mgmt.hybridcontainerservice import HybridContainerServiceMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-hybridcontainerservice
# USAGE
    python put_agent_pool.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = HybridContainerServiceMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.agent_pool.begin_create_or_update(
        connected_cluster_resource_uri="subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/test-hybridakscluster",
        agent_pool_name="testnodepool",
        agent_pool={
            "properties": {
                "count": 1,
                "nodeLabels": {"env": "dev", "goal": "test"},
                "nodeTaints": ["env=prod:NoSchedule", "sku=gpu:NoSchedule"],
                "osType": "Linux",
                "vmSize": "Standard_A4_v2",
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/stable/2024-01-01/examples/PutAgentPool.json
if __name__ == "__main__":
    main()
