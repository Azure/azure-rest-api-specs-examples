from azure.identity import DefaultAzureCredential

from azure.mgmt.devcenter import DevCenterMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-devcenter
# USAGE
    python environment_definitions_list_by_project_catalog.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DevCenterMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="0ac520ee-14c0-480f-b6c9-0a90c58ffff",
    )

    response = client.environment_definitions.list_by_project_catalog(
        resource_group_name="rg1",
        project_name="DevProject",
        catalog_name="myCatalog",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/EnvironmentDefinitions_ListByProjectCatalog.json
if __name__ == "__main__":
    main()
