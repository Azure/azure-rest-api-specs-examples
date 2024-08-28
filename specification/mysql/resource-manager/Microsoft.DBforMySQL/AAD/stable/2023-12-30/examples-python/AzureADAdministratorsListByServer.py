from azure.identity import DefaultAzureCredential

from azure.mgmt.mysqlflexibleservers import MySQLManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-mysqlflexibleservers
# USAGE
    python azure_ad_administrators_list_by_server.py

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

    response = client.azure_ad_administrators.list_by_server(
        resource_group_name="testrg",
        server_name="mysqltestsvc4",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/AAD/stable/2023-12-30/examples/AzureADAdministratorsListByServer.json
if __name__ == "__main__":
    main()
