from azure.identity import DefaultAzureCredential

from azure.mgmt.deviceregistry import DeviceRegistryMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-deviceregistry
# USAGE
    python create_schema_version.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DeviceRegistryMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.schema_versions.create_or_replace(
        resource_group_name="myResourceGroup",
        schema_registry_name="my-schema-registry",
        schema_name="my-schema",
        schema_version_name="1",
        resource={
            "properties": {
                "description": "Schema version 1",
                "schemaContent": '{"$schema": "http://json-schema.org/draft-07/schema#","type": "object","properties": {"humidity": {"type": "string"},"temperature": {"type":"number"}}}',
            }
        },
    )
    print(response)


# x-ms-original-file: 2024-09-01-preview/Create_SchemaVersion.json
if __name__ == "__main__":
    main()
