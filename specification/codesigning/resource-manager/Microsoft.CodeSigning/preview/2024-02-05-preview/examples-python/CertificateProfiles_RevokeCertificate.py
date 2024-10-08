from azure.identity import DefaultAzureCredential

from azure.mgmt.trustedsigning import TrustedSigningMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-trustedsigning
# USAGE
    python certificate_profiles_revoke_certificate.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = TrustedSigningMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    client.certificate_profiles.revoke_certificate(
        resource_group_name="MyResourceGroup",
        account_name="MyAccount",
        profile_name="profileA",
        body={
            "effectiveAt": "2023-11-12T23:40:25+00:00",
            "reason": "KeyCompromised",
            "remarks": "test",
            "serialNumber": "xxxxxxxxxxxxxxxxxx",
            "thumbprint": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
        },
    )


# x-ms-original-file: 2024-02-05-preview/CertificateProfiles_RevokeCertificate.json
if __name__ == "__main__":
    main()
