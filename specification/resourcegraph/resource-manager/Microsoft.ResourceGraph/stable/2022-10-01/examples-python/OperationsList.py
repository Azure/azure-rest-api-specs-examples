from azure.identity import DefaultAzureCredential
from azure.mgmt.resourcegraph import ResourceGraphClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-resourcegraph
# USAGE
    python operations_list.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ResourceGraphClient(
        credential=DefaultAzureCredential(),
    )

    response = client.operations.list()
    for item in response:
        print(item)


# x-ms-original-file: specification/resourcegraph/resource-manager/Microsoft.ResourceGraph/stable/2022-10-01/examples/OperationsList.json
if __name__ == "__main__":
    main()
