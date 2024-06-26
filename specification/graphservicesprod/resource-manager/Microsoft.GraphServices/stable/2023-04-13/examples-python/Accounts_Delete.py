from azure.identity import DefaultAzureCredential
from azure.mgmt.graphservices import GraphServicesMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-graphservices
# USAGE
    python accounts_delete.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = GraphServicesMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="11111111-aaaa-1111-bbbb-111111111111",
    )

    client.accounts.delete(
        resource_group_name="testResourceGroupGRAM",
        resource_name="11111111-aaaa-1111-bbbb-111111111111",
    )


# x-ms-original-file: specification/graphservicesprod/resource-manager/Microsoft.GraphServices/stable/2023-04-13/examples/Accounts_Delete.json
if __name__ == "__main__":
    main()
