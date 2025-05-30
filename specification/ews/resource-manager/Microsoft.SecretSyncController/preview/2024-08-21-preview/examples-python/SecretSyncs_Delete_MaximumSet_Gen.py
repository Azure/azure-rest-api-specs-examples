from azure.identity import DefaultAzureCredential

from azure.mgmt.secretsstoreextension import SecretsStoreExtensionMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-secretsstoreextension
# USAGE
    python secret_syncs_delete_maximum_set_gen.py

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

    client.secret_syncs.begin_delete(
        resource_group_name="rg-ssc-example",
        secret_sync_name="secretsync-ssc-example",
    ).result()


# x-ms-original-file: 2024-08-21-preview/SecretSyncs_Delete_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
