from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.apimanagement import ApiManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-apimanagement
# USAGE
    python api_management_create_documentation.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ApiManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.documentation.create_or_update(
        resource_group_name="rg1",
        service_name="apimService1",
        documentation_id="57d1f7558aa04f15146d9d8a",
        parameters={"properties": {"content": "content", "title": "Title"}},
    )
    print(response)


# x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateDocumentation.json
if __name__ == "__main__":
    main()
