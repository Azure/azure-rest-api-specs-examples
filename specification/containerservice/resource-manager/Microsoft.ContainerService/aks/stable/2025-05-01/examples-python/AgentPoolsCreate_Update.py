from azure.identity import DefaultAzureCredential

from azure.mgmt.containerservice import ContainerServiceClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-containerservice
# USAGE
    python agent_pools_create_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ContainerServiceClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.agent_pools.begin_create_or_update(
        resource_group_name="rg1",
        resource_name="clustername1",
        agent_pool_name="agentpool1",
        parameters={
            "properties": {
                "count": 3,
                "mode": "User",
                "nodeLabels": {"key1": "val1"},
                "nodeTaints": ["Key1=Value1:NoSchedule"],
                "orchestratorVersion": "",
                "osType": "Linux",
                "scaleSetEvictionPolicy": "Delete",
                "scaleSetPriority": "Spot",
                "tags": {"name1": "val1"},
                "vmSize": "Standard_DS1_v2",
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-05-01/examples/AgentPoolsCreate_Update.json
if __name__ == "__main__":
    main()
