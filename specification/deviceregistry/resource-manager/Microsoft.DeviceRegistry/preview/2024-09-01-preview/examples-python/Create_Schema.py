from azure.identity import DefaultAzureCredential

from azure.mgmt.deviceregistry import DeviceRegistryMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-deviceregistry
# USAGE
    python create_schema.py

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

    response = client.schemas.create_or_replace(
        resource_group_name="myResourceGroup",
        schema_registry_name="my-schema-registry",
        schema_name="my-schema",
        resource={
            "properties": {
                "description": "This is a sample Schema",
                "displayName": "My Schema",
                "format": "JsonSchema/draft-07",
                "schemaType": "MessageSchema",
                "tags": {"sampleKey": "sampleValue"},
            }
        },
    )
    print(response)


# x-ms-original-file: 2024-09-01-preview/Create_Schema.json
if __name__ == "__main__":
    main()
