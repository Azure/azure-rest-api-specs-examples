from azure.identity import DefaultAzureCredential
from azure.mgmt.synapse import SynapseManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-synapse
# USAGE
    python kusto_pool_databases_create_or_update.py

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

    response = client.kusto_pool_databases.begin_create_or_update(
        resource_group_name="kustorptest",
        workspace_name="synapseWorkspaceName",
        kusto_pool_name="kustoclusterrptest4",
        database_name="KustoDatabase8",
        parameters={"kind": "ReadWrite", "location": "westus", "properties": {"softDeletePeriod": "P1D"}},
    ).result()
    print(response)


# x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDatabasesCreateOrUpdate.json
if __name__ == "__main__":
    main()
