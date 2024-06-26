from azure.identity import DefaultAzureCredential
from azure.mgmt.synapse import SynapseManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-synapse
# USAGE
    python kusto_pool_database_principal_assignments_check_name_availability.py

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

    response = client.kusto_pool_database_principal_assignments.check_name_availability(
        workspace_name="synapseWorkspaceName",
        kusto_pool_name="kustoclusterrptest4",
        database_name="Kustodatabase8",
        resource_group_name="kustorptest",
        principal_assignment_name={
            "name": "kustoprincipal1",
            "type": "Microsoft.Synapse/workspaces/kustoPools/databases/principalAssignments",
        },
    )
    print(response)


# x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDatabasePrincipalAssignmentsCheckNameAvailability.json
if __name__ == "__main__":
    main()
