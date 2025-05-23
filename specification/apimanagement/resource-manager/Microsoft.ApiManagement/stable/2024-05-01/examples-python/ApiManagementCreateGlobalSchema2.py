from azure.identity import DefaultAzureCredential

from azure.mgmt.apimanagement import ApiManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-apimanagement
# USAGE
    python api_management_create_global_schema2.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ApiManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.global_schema.begin_create_or_update(
        resource_group_name="rg1",
        service_name="apimService1",
        schema_id="schema1",
        parameters={
            "properties": {
                "description": "sample schema description",
                "document": {
                    "$id": "https://example.com/person.schema.json",
                    "$schema": "https://json-schema.org/draft/2020-12/schema",
                    "properties": {
                        "age": {
                            "description": "Age in years which must be equal to or greater than zero.",
                            "minimum": 0,
                            "type": "integer",
                        },
                        "firstName": {"description": "The person's first name.", "type": "string"},
                        "lastName": {"description": "The person's last name.", "type": "string"},
                    },
                    "title": "Person",
                    "type": "object",
                },
                "schemaType": "json",
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateGlobalSchema2.json
if __name__ == "__main__":
    main()
