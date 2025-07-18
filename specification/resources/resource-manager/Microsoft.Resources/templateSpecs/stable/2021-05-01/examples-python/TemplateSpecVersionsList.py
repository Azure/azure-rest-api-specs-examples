from azure.identity import DefaultAzureCredential

from azure.mgmt.resource.templatespecs import TemplateSpecsClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-resource-templatespecs
# USAGE
    python template_spec_versions_list.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = TemplateSpecsClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.template_spec_versions.list(
        resource_group_name="templateSpecRG",
        template_spec_name="simpleTemplateSpec",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/templateSpecs/stable/2021-05-01/examples/TemplateSpecVersionsList.json
if __name__ == "__main__":
    main()
