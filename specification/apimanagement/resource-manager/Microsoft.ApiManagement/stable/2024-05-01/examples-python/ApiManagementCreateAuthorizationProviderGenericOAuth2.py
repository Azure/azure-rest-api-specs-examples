from azure.identity import DefaultAzureCredential

from azure.mgmt.apimanagement import ApiManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-apimanagement
# USAGE
    python api_management_create_authorization_provider_generic_oauth2.py

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

    response = client.authorization_provider.create_or_update(
        resource_group_name="rg1",
        service_name="apimService1",
        authorization_provider_id="eventbrite",
        parameters={
            "properties": {
                "displayName": "eventbrite",
                "identityProvider": "oauth2",
                "oauth2": {
                    "grantTypes": {
                        "authorizationCode": {
                            "authorizationUrl": "https://www.eventbrite.com/oauth/authorize",
                            "clientId": "clientid",
                            "clientSecret": "clientsecretvalue",
                            "refreshUrl": "https://www.eventbrite.com/oauth/token",
                            "scopes": None,
                            "tokenUrl": "https://www.eventbrite.com/oauth/token",
                        }
                    },
                    "redirectUrl": "https://authorization-manager.consent.azure-apim.net/redirect/apim/apimService1",
                },
            }
        },
    )
    print(response)


# x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateAuthorizationProviderGenericOAuth2.json
if __name__ == "__main__":
    main()
