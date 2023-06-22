from azure.identity import DefaultAzureCredential
from azure.mgmt.iothubprovisioningservices import IotDpsClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-iothubprovisioningservices
# USAGE
    python dps_get.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = IotDpsClient(
        credential=DefaultAzureCredential(),
        subscription_id="91d12660-3dec-467a-be2a-213b5544ddc0",
    )

    response = client.iot_dps_resource.get(
        provisioning_service_name="myFirstProvisioningService",
        resource_group_name="myResourceGroup",
    )
    print(response)


# x-ms-original-file: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/preview/2023-03-01-preview/examples/DPSGet.json
if __name__ == "__main__":
    main()
