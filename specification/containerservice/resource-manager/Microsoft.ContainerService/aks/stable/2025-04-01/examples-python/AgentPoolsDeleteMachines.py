from azure.identity import DefaultAzureCredential

from azure.mgmt.containerservice import ContainerServiceClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-containerservice
# USAGE
    python agent_pools_delete_machines.py

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

    client.agent_pools.begin_delete_machines(
        resource_group_name="rg1",
        resource_name="clustername1",
        agent_pool_name="agentpool1",
        machines={"machineNames": ["aks-nodepool1-42263519-vmss00000a", "aks-nodepool1-42263519-vmss00000b"]},
    ).result()


# x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-04-01/examples/AgentPoolsDeleteMachines.json
if __name__ == "__main__":
    main()
