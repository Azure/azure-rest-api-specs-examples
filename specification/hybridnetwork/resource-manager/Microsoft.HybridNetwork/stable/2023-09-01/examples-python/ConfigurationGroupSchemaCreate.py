from azure.identity import DefaultAzureCredential
from azure.mgmt.hybridnetwork import HybridNetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-hybridnetwork
# USAGE
    python configuration_group_schema_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = HybridNetworkManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.configuration_group_schemas.begin_create_or_update(
        resource_group_name="rg1",
        publisher_name="testPublisher",
        configuration_group_schema_name="testConfigurationGroupSchema",
        parameters={
            "location": "westUs2",
            "properties": {
                "description": "Schema with no secrets",
                "schemaDefinition": '{"type":"object","properties":{"interconnect-groups":{"type":"object","properties":{"type":"object","properties":{"name":{"type":"string"},"international-interconnects":{"type":"array","item":{"type":"string"}},"domestic-interconnects":{"type":"array","item":{"type":"string"}}}}},"interconnect-group-assignments":{"type":"object","properties":{"type":"object","properties":{"ssc":{"type":"string"},"interconnects-interconnects":{"type":"string"}}}}},"required":["interconnect-groups","interconnect-group-assignments"]}',
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/ConfigurationGroupSchemaCreate.json
if __name__ == "__main__":
    main()
