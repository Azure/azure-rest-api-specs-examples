from azure.identity import DefaultAzureCredential

from azure.mgmt.apimanagement import ApiManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-apimanagement
# USAGE
    python api_management_create_authorization_server.py

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

    response = client.authorization_server.create_or_update(
        resource_group_name="rg1",
        service_name="apimService1",
        authsid="newauthServer",
        parameters={
            "properties": {
                "authorizationEndpoint": "https://www.contoso.com/oauth2/auth",
                "authorizationMethods": ["GET"],
                "bearerTokenSendingMethods": ["authorizationHeader"],
                "clientId": "1",
                "clientRegistrationEndpoint": "https://www.contoso.com/apps",
                "clientSecret": "2",
                "defaultScope": "read write",
                "description": "test server",
                "displayName": "test2",
                "grantTypes": ["authorizationCode", "implicit"],
                "resourceOwnerPassword": "pwd",
                "resourceOwnerUsername": "un",
                "supportState": True,
                "tokenEndpoint": "https://www.contoso.com/oauth2/token",
                "useInApiDocumentation": True,
                "useInTestConsole": False,
            }
        },
    )
    print(response)


# x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateAuthorizationServer.json
if __name__ == "__main__":
    main()
