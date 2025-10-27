from azure.identity import DefaultAzureCredential

from azure.mgmt.deviceregistry import DeviceRegistryMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-deviceregistry
# USAGE
    python list_schemas_by_schema_registry.py

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

    response = client.schemas.list_by_schema_registry(
        resource_group_name="myResourceGroup",
        schema_registry_name="my-schema-registry",
    )
    for item in response:
        print(item)


# x-ms-original-file: 2025-10-01/List_Schemas_BySchemaRegistry.json
if __name__ == "__main__":
    main()
