from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.rdbms.postgresql import PostgreSQLManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-rdbms
# USAGE
    python server_create_point_in_time_restore.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = PostgreSQLManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="ffffffff-ffff-ffff-ffff-ffffffffffff",
    )

    response = client.servers.begin_create(
        resource_group_name="TargetResourceGroup",
        server_name="targetserver",
        parameters={
            "location": "brazilsouth",
            "properties": {
                "createMode": "PointInTimeRestore",
                "restorePointInTime": "2017-12-14T00:00:37.467Z",
                "sourceServerId": "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/SourceResourceGroup/providers/Microsoft.DBforPostgreSQL/servers/sourceserver",
            },
            "sku": {"capacity": 2, "family": "Gen5", "name": "B_Gen5_2", "tier": "Basic"},
            "tags": {"ElasticServer": "1"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/ServerCreatePointInTimeRestore.json
if __name__ == "__main__":
    main()
