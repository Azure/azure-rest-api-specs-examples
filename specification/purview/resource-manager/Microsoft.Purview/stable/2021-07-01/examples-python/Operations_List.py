from azure.identity import DefaultAzureCredential
from azure.mgmt.purview import PurviewManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-purview
# USAGE
    python operations_list.py

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

    response = client.operations.list()
    for item in response:
        print(item)


# x-ms-original-file: specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/Operations_List.json
if __name__ == "__main__":
    main()
