from azure.identity import DefaultAzureCredential
from azure.mgmt.attestation import AttestationManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-attestation
# USAGE
    python attestation_providers_get.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AttestationManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.attestation_providers.get(
        resource_group_name="MyResourceGroup",
        provider_name="myattestationprovider",
    )
    print(response)


# x-ms-original-file: specification/attestation/resource-manager/Microsoft.Attestation/stable/2020-10-01/examples/Get_AttestationProvider.json
if __name__ == "__main__":
    main()
