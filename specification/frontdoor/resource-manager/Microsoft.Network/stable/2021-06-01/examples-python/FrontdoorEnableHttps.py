from azure.identity import DefaultAzureCredential
from azure.mgmt.frontdoor import FrontDoorManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-frontdoor
# USAGE
    python frontdoor_enable_https.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = FrontDoorManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.frontend_endpoints.begin_enable_https(
        resource_group_name="rg1",
        front_door_name="frontDoor1",
        frontend_endpoint_name="frontendEndpoint1",
        custom_https_configuration={
            "certificateSource": "AzureKeyVault",
            "keyVaultCertificateSourceParameters": {
                "secretName": "secret1",
                "secretVersion": "00000000-0000-0000-0000-000000000000",
                "vault": {"id": "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.KeyVault/vaults/vault1"},
            },
            "minimumTlsVersion": "1.0",
            "protocolType": "ServerNameIndication",
        },
    ).result()
    print(response)


# x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2021-06-01/examples/FrontdoorEnableHttps.json
if __name__ == "__main__":
    main()
