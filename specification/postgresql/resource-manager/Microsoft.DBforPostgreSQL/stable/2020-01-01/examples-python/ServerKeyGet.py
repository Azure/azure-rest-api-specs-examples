from azure.identity import DefaultAzureCredential

from azure.mgmt.rdbms.postgresql import PostgreSQLManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-rdbms
# USAGE
    python server_key_get.py

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

    response = client.server_keys.get(
        resource_group_name="testrg",
        server_name="testserver",
        key_name="someVault_someKey_01234567890123456789012345678901",
    )
    print(response)


# x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2020-01-01/examples/ServerKeyGet.json
if __name__ == "__main__":
    main()
