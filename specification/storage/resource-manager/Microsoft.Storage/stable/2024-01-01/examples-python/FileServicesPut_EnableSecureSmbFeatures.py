from azure.identity import DefaultAzureCredential

from azure.mgmt.storage import StorageManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-storage
# USAGE
    python file_services_put_enable_secure_smb_features.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = StorageManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="{subscription-id}",
    )

    response = client.file_services.set_service_properties(
        resource_group_name="res4410",
        account_name="sto8607",
        parameters={
            "properties": {
                "protocolSettings": {
                    "smb": {
                        "authenticationMethods": "NTLMv2;Kerberos",
                        "channelEncryption": "AES-128-CCM;AES-128-GCM;AES-256-GCM",
                        "kerberosTicketEncryption": "RC4-HMAC;AES-256",
                        "versions": "SMB2.1;SMB3.0;SMB3.1.1",
                    }
                }
            }
        },
    )
    print(response)


# x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/FileServicesPut_EnableSecureSmbFeatures.json
if __name__ == "__main__":
    main()
