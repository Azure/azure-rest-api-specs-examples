from azure.identity import DefaultAzureCredential
from azure.mgmt.synapse import SynapseManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-synapse
# USAGE
    python kusto_pools_start.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SynapseManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="12345678-1234-1234-1234-123456789098",
    )

    response = client.kusto_pools.begin_start(
        workspace_name="kustorptest",
        kusto_pool_name="kustoclusterrptest4",
        resource_group_name="kustorptest",
    ).result()
    print(response)


# x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsStart.json
if __name__ == "__main__":
    main()
