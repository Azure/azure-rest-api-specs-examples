from azure.identity import DefaultAzureCredential

from azure.mgmt.msi import ManagedServiceIdentityClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-msi
# USAGE
    python identity_delete.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ManagedServiceIdentityClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    client.user_assigned_identities.delete(
        resource_group_name="rgName",
        resource_name="resourceName",
    )


# x-ms-original-file: specification/msi/resource-manager/Microsoft.ManagedIdentity/stable/2024-11-30/examples/IdentityDelete.json
if __name__ == "__main__":
    main()
