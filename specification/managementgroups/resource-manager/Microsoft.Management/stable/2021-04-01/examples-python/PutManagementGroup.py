from azure.identity import DefaultAzureCredential

from azure.mgmt.managementgroups import ManagementGroupsAPI

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-managementgroups
# USAGE
    python put_management_group.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ManagementGroupsAPI(
        credential=DefaultAzureCredential(),
    )

    response = client.management_groups.begin_create_or_update(
        group_id="ChildGroup",
        create_management_group_request={
            "properties": {
                "details": {"parent": {"id": "/providers/Microsoft.Management/managementGroups/RootGroup"}},
                "displayName": "ChildGroup",
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/PutManagementGroup.json
if __name__ == "__main__":
    main()
