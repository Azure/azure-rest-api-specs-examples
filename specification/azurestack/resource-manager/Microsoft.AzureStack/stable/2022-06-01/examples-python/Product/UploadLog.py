from azure.identity import DefaultAzureCredential
from azure.mgmt.azurestack import AzureStackManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-azurestack
# USAGE
    python upload_log.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AzureStackManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="dd8597b4-8739-4467-8b10-f8679f62bfbf",
    )

    response = client.products.upload_log(
        resource_group="azurestack",
        registration_name="testregistration",
        product_name="Microsoft.OSTCExtensions.VMAccessForLinux.1.4.7.1",
    )
    print(response)


# x-ms-original-file: specification/azurestack/resource-manager/Microsoft.AzureStack/stable/2022-06-01/examples/Product/UploadLog.json
if __name__ == "__main__":
    main()
