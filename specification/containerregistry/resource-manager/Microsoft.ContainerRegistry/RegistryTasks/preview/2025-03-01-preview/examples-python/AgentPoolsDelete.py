from azure.identity import DefaultAzureCredential

from azure.mgmt.containerregistrytasks import ContainerRegistryTasksMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-containerregistrytasks
# USAGE
    python agent_pools_delete.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ContainerRegistryTasksMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    client.agent_pools.begin_delete(
        resource_group_name="myResourceGroup",
        registry_name="myRegistry",
        agent_pool_name="myAgentPool",
    ).result()


# x-ms-original-file: 2025-03-01-preview/AgentPoolsDelete.json
if __name__ == "__main__":
    main()
