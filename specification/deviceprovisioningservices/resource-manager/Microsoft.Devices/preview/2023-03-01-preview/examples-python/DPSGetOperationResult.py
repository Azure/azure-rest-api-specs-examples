from azure.identity import DefaultAzureCredential
from azure.mgmt.iothubprovisioningservices import IotDpsClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-iothubprovisioningservices
# USAGE
    python dps_get_operation_result.py

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

    response = client.iot_dps_resource.get_operation_result(
        operation_id="MTY5OTNmZDctODI5Yy00N2E2LTkxNDQtMDU1NGIyYzY1ZjRl",
        resource_group_name="myResourceGroup",
        provisioning_service_name="myFirstProvisioningService",
    )
    print(response)


# x-ms-original-file: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/preview/2023-03-01-preview/examples/DPSGetOperationResult.json
if __name__ == "__main__":
    main()
