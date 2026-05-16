from azure.identity import DefaultAzureCredential

from azure.mgmt.azurestack import AzureStackManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-azurestack
# USAGE
    python list.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AzureStackManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.cloud_manifest_file.list()
    print(response)


# x-ms-original-file: specification/azurestack/resource-manager/Microsoft.AzureStack/AzureStack/preview/2020-06-01-preview/examples/CloudManifestFile/List.json
if __name__ == "__main__":
    main()
