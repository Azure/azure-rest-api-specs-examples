from azure.identity import DefaultAzureCredential

from azure.mgmt.rdbms.mariadb import MariaDBManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-rdbms
# USAGE
    python top_query_statistics_get.py

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

    response = client.top_query_statistics.get(
        resource_group_name="testResourceGroupName",
        server_name="testServerName",
        query_statistic_id="66-636923268000000000-636923277000000000-avg-duration",
    )
    print(response)


# x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/TopQueryStatisticsGet.json
if __name__ == "__main__":
    main()
