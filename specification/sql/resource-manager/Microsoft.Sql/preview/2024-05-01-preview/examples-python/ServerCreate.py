from azure.identity import DefaultAzureCredential

from azure.mgmt.sql import SqlManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-sql
# USAGE
    python server_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SqlManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-1111-2222-3333-444444444444",
    )

    response = client.servers.begin_create_or_update(
        resource_group_name="sqlcrudtest-7398",
        server_name="sqlcrudtest-4645",
        parameters={
            "location": "Japan East",
            "properties": {
                "administratorLogin": "dummylogin",
                "administratorLoginPassword": "PLACEHOLDER",
                "administrators": {
                    "azureADOnlyAuthentication": True,
                    "login": "bob@contoso.com",
                    "principalType": "User",
                    "sid": "00000011-1111-2222-2222-123456789111",
                    "tenantId": "00000011-1111-2222-2222-123456789111",
                },
                "isIPv6Enabled": "Enabled",
                "publicNetworkAccess": "Enabled",
                "restrictOutboundNetworkAccess": "Enabled",
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2024-05-01-preview/examples/ServerCreate.json
if __name__ == "__main__":
    main()
