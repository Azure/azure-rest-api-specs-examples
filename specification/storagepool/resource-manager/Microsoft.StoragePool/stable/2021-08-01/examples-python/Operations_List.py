from azure.identity import DefaultAzureCredential
from azure.mgmt.storagepool import StoragePoolManagement

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-storagepool
# USAGE
    python list_operations.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = StoragePoolManagement(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.operations.list()
    for item in response:
        print(item)


# x-ms-original-file: specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/Operations_List.json
if __name__ == "__main__":
    main()
