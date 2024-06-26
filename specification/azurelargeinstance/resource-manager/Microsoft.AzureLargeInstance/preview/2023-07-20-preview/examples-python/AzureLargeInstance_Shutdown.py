from azure.identity import DefaultAzureCredential
from azure.mgmt.azurelargeinstance import LargeInstanceMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-azurelargeinstance
# USAGE
    python azure_large_instance_shutdown.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = LargeInstanceMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.azure_large_instance.begin_shutdown(
        resource_group_name="myResourceGroup",
        azure_large_instance_name="myALInstance",
    ).result()
    print(response)


# x-ms-original-file: specification/azurelargeinstance/resource-manager/Microsoft.AzureLargeInstance/preview/2023-07-20-preview/examples/AzureLargeInstance_Shutdown.json
if __name__ == "__main__":
    main()
