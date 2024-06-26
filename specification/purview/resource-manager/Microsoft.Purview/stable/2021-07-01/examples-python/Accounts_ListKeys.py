from azure.identity import DefaultAzureCredential
from azure.mgmt.purview import PurviewManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-purview
# USAGE
    python accounts_list_keys.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = PurviewManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="12345678-1234-1234-12345678abc",
    )

    response = client.accounts.list_keys(
        resource_group_name="SampleResourceGroup",
        account_name="account1",
    )
    print(response)


# x-ms-original-file: specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/Accounts_ListKeys.json
if __name__ == "__main__":
    main()
