from azure.identity import DefaultAzureCredential
from azure.mgmt.datashare import DataShareManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-datashare
# USAGE
    python data_set_mappings_sql_db_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DataShareManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="433a8dfd-e5d5-4e77-ad86-90acdc75eb1a",
    )

    response = client.data_set_mappings.create(
        resource_group_name="SampleResourceGroup",
        account_name="Account1",
        share_subscription_name="ShareSubscription1",
        data_set_mapping_name="DatasetMapping1",
        data_set_mapping={
            "kind": "SqlDBTable",
            "properties": {
                "dataSetId": "a08f184b-0567-4b11-ba22-a1199336d226",
                "databaseName": "Database1",
                "schemaName": "dbo",
                "sqlServerResourceId": "/subscriptions/433a8dfd-e5d5-4e77-ad86-90acdc75eb1a/resourceGroups/SampleResourceGroup/providers/Microsoft.Sql/servers/Server1",
                "tableName": "Table1",
            },
        },
    )
    print(response)


# x-ms-original-file: specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/DataSetMappings_SqlDB_Create.json
if __name__ == "__main__":
    main()
