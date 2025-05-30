from azure.identity import DefaultAzureCredential

from azure.mgmt.apimanagement import ApiManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-apimanagement
# USAGE
    python api_management_create_multi_region_service_with_custom_hostname.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ApiManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.api_management_service.begin_create_or_update(
        resource_group_name="rg1",
        service_name="apimService1",
        parameters={
            "location": "West US",
            "properties": {
                "additionalLocations": [
                    {"disableGateway": True, "location": "East US", "sku": {"capacity": 1, "name": "Premium"}}
                ],
                "apiVersionConstraint": {"minApiVersion": "2019-01-01"},
                "hostnameConfigurations": [
                    {
                        "certificatePassword": "Password",
                        "defaultSslBinding": True,
                        "encodedCertificate": "****** Base 64 Encoded Certificate ************",
                        "hostName": "gateway1.msitesting.net",
                        "type": "Proxy",
                    },
                    {
                        "certificatePassword": "Password",
                        "encodedCertificate": "****** Base 64 Encoded Certificate ************",
                        "hostName": "mgmt.msitesting.net",
                        "type": "Management",
                    },
                    {
                        "certificatePassword": "Password",
                        "encodedCertificate": "****** Base 64 Encoded Certificate ************",
                        "hostName": "portal1.msitesting.net",
                        "type": "Portal",
                    },
                    {
                        "certificatePassword": "Password",
                        "encodedCertificate": "****** Base 64 Encoded Certificate ************",
                        "hostName": "configuration-api.msitesting.net",
                        "type": "ConfigurationApi",
                    },
                ],
                "publisherEmail": "apim@autorestsdk.com",
                "publisherName": "autorestsdk",
                "virtualNetworkType": "None",
            },
            "sku": {"capacity": 1, "name": "Premium"},
            "tags": {"tag1": "value1", "tag2": "value2", "tag3": "value3"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateMultiRegionServiceWithCustomHostname.json
if __name__ == "__main__":
    main()
