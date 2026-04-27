from azure.identity import DefaultAzureCredential

from azure.mgmt.confluent import ConfluentManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-confluent
# USAGE
    python access_delete_role_binding_minimum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ConfluentManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    client.access.delete_role_binding(
        resource_group_name="rgconfluent",
        organization_name="kxbwvfhsqesuaswozdiivwo",
        role_binding_id="dqlmrdp",
    )


# x-ms-original-file: 2025-08-18-preview/Access_DeleteRoleBinding_MinimumSet_Gen.json
if __name__ == "__main__":
    main()
