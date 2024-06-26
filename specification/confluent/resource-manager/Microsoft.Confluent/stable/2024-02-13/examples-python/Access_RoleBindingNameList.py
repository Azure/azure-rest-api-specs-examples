from azure.identity import DefaultAzureCredential
from azure.mgmt.confluent import ConfluentManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-confluent
# USAGE
    python access_role_binding_name_list.py

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

    response = client.access.list_role_binding_name_list(
        resource_group_name="myResourceGroup",
        organization_name="myOrganization",
        body={
            "searchFilters": {
                "crn_pattern": "crn://confluent.cloud/organization=1aa7de07-298e-479c-8f2f-16ac91fd8e76",
                "namespace": "public,dataplane,networking,identity,datagovernance,connect,streamcatalog,pipelines,ksql",
            }
        },
    )
    print(response)


# x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/Access_RoleBindingNameList.json
if __name__ == "__main__":
    main()
