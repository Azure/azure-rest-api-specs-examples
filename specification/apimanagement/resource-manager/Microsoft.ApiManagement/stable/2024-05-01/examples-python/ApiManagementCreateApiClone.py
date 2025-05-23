from azure.identity import DefaultAzureCredential

from azure.mgmt.apimanagement import ApiManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-apimanagement
# USAGE
    python api_management_create_api_clone.py

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

    response = client.api.begin_create_or_update(
        resource_group_name="rg1",
        service_name="apimService1",
        api_id="echo-api2",
        parameters={
            "properties": {
                "description": "Copy of Existing Echo Api including Operations.",
                "displayName": "Echo API2",
                "isCurrent": True,
                "path": "echo2",
                "protocols": ["http", "https"],
                "serviceUrl": "http://echoapi.cloudapp.net/api",
                "sourceApiId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/58a4aeac497000007d040001",
                "subscriptionRequired": True,
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateApiClone.json
if __name__ == "__main__":
    main()
