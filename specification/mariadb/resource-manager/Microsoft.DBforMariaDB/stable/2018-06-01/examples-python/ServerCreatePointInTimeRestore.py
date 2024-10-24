from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.rdbms.mariadb import MariaDBManagementClient

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
    client = MariaDBManagementClient(
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
                "sourceServerId": "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/SourceResourceGroup/providers/Microsoft.DBforMariaDB/servers/sourceserver",
            },
            "sku": {"capacity": 2, "family": "Gen5", "name": "GP_Gen5_2", "tier": "GeneralPurpose"},
            "tags": {"ElasticServer": "1"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/ServerCreatePointInTimeRestore.json
if __name__ == "__main__":
    main()
