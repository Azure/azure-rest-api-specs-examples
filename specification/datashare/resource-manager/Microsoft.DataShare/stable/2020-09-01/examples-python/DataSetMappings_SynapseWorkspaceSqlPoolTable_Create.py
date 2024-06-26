from azure.identity import DefaultAzureCredential
from azure.mgmt.datashare import DataShareManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-datashare
# USAGE
    python data_set_mappings_synapse_workspace_sql_pool_table_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DataShareManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="0f3dcfc3-18f8-4099-b381-8353e19d43a7",
    )

    response = client.data_set_mappings.create(
        resource_group_name="SampleResourceGroup",
        account_name="consumerAccount",
        share_subscription_name="ShareSubscription1",
        data_set_mapping_name="datasetMappingName1",
        data_set_mapping={
            "kind": "SynapseWorkspaceSqlPoolTable",
            "properties": {
                "dataSetId": "3dc64e49-1fc3-4186-b3dc-d388c4d3076a",
                "synapseWorkspaceSqlPoolTableResourceId": "/subscriptions/0f3dcfc3-18f8-4099-b381-8353e19d43a7/resourceGroups/SampleResourceGroup/providers/Microsoft.Synapse/workspaces/ExampleWorkspace/sqlPools/ExampleSqlPool/schemas/dbo/tables/table1",
            },
        },
    )
    print(response)


# x-ms-original-file: specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/DataSetMappings_SynapseWorkspaceSqlPoolTable_Create.json
if __name__ == "__main__":
    main()
