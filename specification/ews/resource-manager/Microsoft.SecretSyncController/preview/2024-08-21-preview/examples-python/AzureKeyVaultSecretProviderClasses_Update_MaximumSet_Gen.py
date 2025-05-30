from azure.identity import DefaultAzureCredential

from azure.mgmt.secretsstoreextension import SecretsStoreExtensionMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-secretsstoreextension
# USAGE
    python azure_key_vault_secret_provider_classes_update_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SecretsStoreExtensionMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.azure_key_vault_secret_provider_classes.begin_update(
        resource_group_name="rg-ssc-example",
        azure_key_vault_secret_provider_class_name="akvspc-ssc-example",
        properties={
            "properties": {
                "clientId": "00000000-0000-0000-0000-000000000000",
                "keyvaultName": "example-ssc-key-vault",
                "objects": "array: |\n  - |\n    objectName: my-secret-object\n    objectType: secret\n    objectVersionHistory: 1",
                "tenantId": "00000000-0000-0000-0000-000000000000",
            },
            "tags": {"example-tag": "example-tag-value"},
        },
    ).result()
    print(response)


# x-ms-original-file: 2024-08-21-preview/AzureKeyVaultSecretProviderClasses_Update_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
