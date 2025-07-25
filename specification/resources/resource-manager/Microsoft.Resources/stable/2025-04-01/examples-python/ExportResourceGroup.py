from azure.identity import DefaultAzureCredential

from azure.mgmt.resource.resources import ResourceManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-resource
# USAGE
    python export_resource_group.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ResourceManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.resource_groups.begin_export_template(
        resource_group_name="my-resource-group",
        parameters={"options": "IncludeParameterDefaultValue,IncludeComments", "resources": ["*"]},
    ).result()
    print(response)


# x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2025-04-01/examples/ExportResourceGroup.json
if __name__ == "__main__":
    main()
