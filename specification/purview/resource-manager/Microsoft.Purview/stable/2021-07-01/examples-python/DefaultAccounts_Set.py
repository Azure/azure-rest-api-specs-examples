from azure.identity import DefaultAzureCredential
from azure.mgmt.purview import PurviewManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-purview
# USAGE
    python default_accounts_set.py

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

    response = client.default_accounts.set(
        default_account_payload={
            "accountName": "myDefaultAccount",
            "resourceGroupName": "rg-1",
            "scope": "11733A4E-BA84-46FF-91D1-AFF1A3215A90",
            "scopeTenantId": "11733A4E-BA84-46FF-91D1-AFF1A3215A90",
            "scopeType": "Tenant",
            "subscriptionId": "12345678-1234-1234-12345678aaa",
        },
    )
    print(response)


# x-ms-original-file: specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/DefaultAccounts_Set.json
if __name__ == "__main__":
    main()
