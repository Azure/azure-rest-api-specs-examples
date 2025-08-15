from azure.identity import DefaultAzureCredential

from azure.mgmt.iothub import IotHubClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-iothub
# USAGE
    python iothub_listbyrg.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = IotHubClient(
        credential=DefaultAzureCredential(),
        subscription_id="91d12660-3dec-467a-be2a-213b5544ddc0",
    )

    response = client.iot_hub_resource.list_by_resource_group(
        resource_group_name="myResourceGroup",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/preview/2023-06-30-preview/examples/iothub_listbyrg.json
if __name__ == "__main__":
    main()
