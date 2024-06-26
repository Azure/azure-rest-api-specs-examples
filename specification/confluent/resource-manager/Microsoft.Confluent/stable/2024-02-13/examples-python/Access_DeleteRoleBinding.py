from azure.identity import DefaultAzureCredential
from azure.mgmt.confluent import ConfluentManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-confluent
# USAGE
    python access_delete_role_binding.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ConfluentManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    client.access.delete_role_binding(
        resource_group_name="myResourceGroup",
        organization_name="myOrganization",
        role_binding_id="dlz-f3a90de",
    )


# x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/Access_DeleteRoleBinding.json
if __name__ == "__main__":
    main()
