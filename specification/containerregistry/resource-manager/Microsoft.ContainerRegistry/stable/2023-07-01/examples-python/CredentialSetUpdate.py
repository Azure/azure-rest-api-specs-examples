from azure.identity import DefaultAzureCredential

from azure.mgmt.containerregistry import ContainerRegistryManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-containerregistry
# USAGE
    python credential_set_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ContainerRegistryManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.credential_sets.begin_update(
        resource_group_name="myResourceGroup",
        registry_name="myRegistry",
        credential_set_name="myCredentialSet",
        credential_set_update_parameters={
            "properties": {
                "authCredentials": [
                    {
                        "name": "Credential1",
                        "passwordSecretIdentifier": "https://myvault.vault.azure.net/secrets/password2",
                        "usernameSecretIdentifier": "https://myvault.vault.azure.net/secrets/username2",
                    }
                ]
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2023-07-01/examples/CredentialSetUpdate.json
if __name__ == "__main__":
    main()
