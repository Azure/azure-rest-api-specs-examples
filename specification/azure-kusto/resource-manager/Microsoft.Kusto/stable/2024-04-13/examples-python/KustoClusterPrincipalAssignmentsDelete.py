from azure.identity import DefaultAzureCredential

from azure.mgmt.kusto import KustoManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-kusto
# USAGE
    python kusto_cluster_principal_assignments_delete.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = KustoManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="12345678-1234-1234-1234-123456789098",
    )

    client.cluster_principal_assignments.begin_delete(
        resource_group_name="kustorptest",
        cluster_name="kustoCluster",
        principal_assignment_name="kustoprincipal1",
    ).result()


# x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/KustoClusterPrincipalAssignmentsDelete.json
if __name__ == "__main__":
    main()
