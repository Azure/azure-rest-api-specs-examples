from azure.identity import DefaultAzureCredential

from azure.mgmt.iothub import IotHubClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-iothub
# USAGE
    python iot_hub_manual_failover.py

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

    client.iot_hub.begin_manual_failover(
        iot_hub_name="testHub",
        resource_group_name="myResourceGroup",
        failover_input={"failoverRegion": "testHub"},
    ).result()


# x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2023-06-30/examples/IotHub_ManualFailover.json
if __name__ == "__main__":
    main()
