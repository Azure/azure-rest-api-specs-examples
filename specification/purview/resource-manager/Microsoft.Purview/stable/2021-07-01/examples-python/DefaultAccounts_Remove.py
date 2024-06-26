from azure.identity import DefaultAzureCredential
from azure.mgmt.purview import PurviewManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-purview
# USAGE
    python default_accounts_remove.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = PurviewManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.default_accounts.remove(
        scope_tenant_id="11733A4E-BA84-46FF-91D1-AFF1A3215A90",
        scope_type="Tenant",
    )
    print(response)


# x-ms-original-file: specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/DefaultAccounts_Remove.json
if __name__ == "__main__":
    main()
