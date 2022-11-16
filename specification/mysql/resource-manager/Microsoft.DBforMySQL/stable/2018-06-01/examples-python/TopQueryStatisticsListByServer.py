from azure.identity import DefaultAzureCredential
from azure.mgmt.rdbms import MySQLManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-rdbms
# USAGE
    python top_query_statistics_list_by_server.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = MySQLManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="ffffffff-ffff-ffff-ffff-ffffffffffff",
    )

    response = client.top_query_statistics.list_by_server(
        resource_group_name="testResourceGroupName",
        server_name="testServerName",
        parameters={
            "properties": {
                "aggregationFunction": "avg",
                "aggregationWindow": "PT15M",
                "numberOfTopQueries": 5,
                "observationEndTime": "2019-05-07T20:00:00.000Z",
                "observationStartTime": "2019-05-01T20:00:00.000Z",
                "observedMetric": "duration",
            }
        },
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/TopQueryStatisticsListByServer.json
if __name__ == "__main__":
    main()
