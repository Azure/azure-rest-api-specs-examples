from azure.identity import DefaultAzureCredential
from azure.mgmt.azurearcdata import AzureArcDataManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-azurearcdata
# USAGE
    python create_or_update_sql_server_instance.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AzureArcDataManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-1111-2222-3333-444444444444",
    )

    response = client.sql_server_instances.begin_create(
        resource_group_name="testrg",
        sql_server_instance_name="testsqlServerInstance",
        sql_server_instance={
            "location": "northeurope",
            "properties": {
                "azureDefenderStatus": "Protected",
                "azureDefenderStatusLastUpdated": "2020-01-02T17:18:19.1234567Z",
                "collation": "collation",
                "containerResourceId": "Resource id of hosting Arc Machine",
                "currentVersion": "2012",
                "edition": "Developer",
                "hostType": "Physical Server",
                "instanceName": "name of instance",
                "licenseType": "Free",
                "patchLevel": "patchlevel",
                "productId": "sql id",
                "status": "Registered",
                "tcpDynamicPorts": "1433",
                "tcpStaticPorts": "1433",
                "vCore": "4",
                "version": "SQL Server 2012",
            },
            "tags": {"mytag": "myval"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/preview/2022-03-01-preview/examples/CreateOrUpdateSqlServerInstance.json
if __name__ == "__main__":
    main()
