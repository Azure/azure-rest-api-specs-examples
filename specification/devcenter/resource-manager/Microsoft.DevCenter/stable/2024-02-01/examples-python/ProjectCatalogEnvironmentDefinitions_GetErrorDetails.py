from azure.identity import DefaultAzureCredential

from azure.mgmt.devcenter import DevCenterMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-devcenter
# USAGE
    python project_catalog_environment_definitions_get_error_details.py

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

    response = client.project_catalog_environment_definitions.get_error_details(
        resource_group_name="rg1",
        project_name="DevProject",
        catalog_name="myCatalog",
        environment_definition_name="myEnvironmentDefinition",
    )
    print(response)


# x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/ProjectCatalogEnvironmentDefinitions_GetErrorDetails.json
if __name__ == "__main__":
    main()
