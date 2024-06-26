from azure.identity import DefaultAzureCredential
from azure.mgmt.confluent import ConfluentManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-confluent
# USAGE
    python access_create_role_binding.py

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

    response = client.access.create_role_binding(
        resource_group_name="myResourceGroup",
        organization_name="myOrganization",
        body={
            "crn_pattern": "crn://confluent.cloud/organization=1111aaaa-11aa-11aa-11aa-111111aaaaaa/environment=env-aaa1111/cloud-cluster=lkc-1111aaa",
            "principal": "User:u-111aaa",
            "role_name": "CloudClusterAdmin",
        },
    )
    print(response)


# x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/Access_CreateRoleBinding.json
if __name__ == "__main__":
    main()
