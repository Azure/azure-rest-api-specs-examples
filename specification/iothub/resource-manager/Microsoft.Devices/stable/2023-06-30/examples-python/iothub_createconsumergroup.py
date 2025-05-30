from azure.identity import DefaultAzureCredential

from azure.mgmt.iothub import IotHubClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-iothub
# USAGE
    python iothub_createconsumergroup.py

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

    response = client.iot_hub_resource.create_event_hub_consumer_group(
        resource_group_name="myResourceGroup",
        resource_name="testHub",
        event_hub_endpoint_name="events",
        name="test",
        consumer_group_body={"properties": {"name": "test"}},
    )
    print(response)


# x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2023-06-30/examples/iothub_createconsumergroup.json
if __name__ == "__main__":
    main()
