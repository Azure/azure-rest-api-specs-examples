from azure.identity import DefaultAzureCredential

from azure.mgmt.kusto import KustoManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-kusto
# USAGE
    python kusto_attached_database_configurations_create_or_update.py

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

    response = client.attached_database_configurations.begin_create_or_update(
        resource_group_name="kustorptest",
        cluster_name="kustoCluster2",
        attached_database_configuration_name="attachedDatabaseConfigurationsTest",
        parameters={
            "location": "westus",
            "properties": {
                "clusterResourceId": "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster2",
                "databaseName": "kustodatabase",
                "databaseNameOverride": "overridekustodatabase",
                "defaultPrincipalsModificationKind": "Union",
                "tableLevelSharingProperties": {
                    "externalTablesToExclude": ["ExternalTable2"],
                    "externalTablesToInclude": ["ExternalTable1"],
                    "materializedViewsToExclude": ["MaterializedViewTable2"],
                    "materializedViewsToInclude": ["MaterializedViewTable1"],
                    "tablesToExclude": ["Table2"],
                    "tablesToInclude": ["Table1"],
                },
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/KustoAttachedDatabaseConfigurationsCreateOrUpdate.json
if __name__ == "__main__":
    main()
