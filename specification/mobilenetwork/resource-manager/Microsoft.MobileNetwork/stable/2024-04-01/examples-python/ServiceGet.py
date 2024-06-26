from azure.identity import DefaultAzureCredential

from azure.mgmt.mobilenetwork import MobileNetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-mobilenetwork
# USAGE
    python service_get.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = MobileNetworkManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.services.get(
        resource_group_name="rg1",
        mobile_network_name="testMobileNetwork",
        service_name="TestService",
    )
    print(response)


# x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-04-01/examples/ServiceGet.json
if __name__ == "__main__":
    main()
