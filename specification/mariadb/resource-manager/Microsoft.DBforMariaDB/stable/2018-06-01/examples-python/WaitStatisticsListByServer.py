from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.rdbms.mariadb import MariaDBManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-rdbms
# USAGE
    python wait_statistics_list_by_server.py

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

    response = client.wait_statistics.list_by_server(
        resource_group_name="testResourceGroupName",
        server_name="testServerName",
        parameters={
            "properties": {
                "aggregationWindow": "PT15M",
                "observationEndTime": "2019-05-07T20:00:00.000Z",
                "observationStartTime": "2019-05-01T20:00:00.000Z",
            }
        },
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/WaitStatisticsListByServer.json
if __name__ == "__main__":
    main()
