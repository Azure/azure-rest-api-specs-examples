from azure.identity import DefaultAzureCredential

from azure.mgmt.managementgroups import ManagementGroupsAPI

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-managementgroups
# USAGE
    python get_management_group_with_expand.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ManagementGroupsAPI(
        credential=DefaultAzureCredential(),
    )

    response = client.management_groups.get(
        group_id="20000000-0001-0000-0000-000000000000",
    )
    print(response)


# x-ms-original-file: specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/GetManagementGroupWithExpand.json
if __name__ == "__main__":
    main()
