
from azure.identity import DefaultAzureCredential
from azure.mgmt.cdn import CdnManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-cdn
# USAGE
    python validate_secret.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret 
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = CdnManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.validate.secret(
        validate_secret_input={
            "secretSource": {
                "id": "/subscriptions/subid/resourcegroups/RG/providers/Microsoft.KeyVault/vault/kvName/certificate/certName"
            },
            "secretType": "CustomerCertificate",
        },
    )
    print(response)


# x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Validate_Secret.json
if __name__ == "__main__":
    main()
