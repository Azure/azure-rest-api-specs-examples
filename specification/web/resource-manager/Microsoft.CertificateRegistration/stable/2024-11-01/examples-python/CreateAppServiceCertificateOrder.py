from azure.identity import DefaultAzureCredential

from azure.mgmt.web import WebSiteManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-web
# USAGE
    python create_app_service_certificate_order.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = WebSiteManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="34adfa4f-cedf-4dc0-ba29-b6d1a69ab345",
    )

    response = client.app_service_certificate_orders.begin_create_or_update(
        resource_group_name="testrg123",
        certificate_order_name="SampleCertificateOrderName",
        certificate_distinguished_name={
            "location": "Global",
            "properties": {
                "autoRenew": True,
                "certificates": {
                    "SampleCertName1": {
                        "keyVaultId": "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/testrg123/providers/microsoft.keyvault/vaults/SamplevaultName",
                        "keyVaultSecretName": "SampleSecretName1",
                    },
                    "SampleCertName2": {
                        "keyVaultId": "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/testrg123/providers/microsoft.keyvault/vaults/SamplevaultName",
                        "keyVaultSecretName": "SampleSecretName2",
                    },
                },
                "distinguishedName": "CN=SampleCustomDomain.com",
                "keySize": 2048,
                "productType": "StandardDomainValidatedSsl",
                "validityInYears": 2,
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2024-11-01/examples/CreateAppServiceCertificateOrder.json
if __name__ == "__main__":
    main()
