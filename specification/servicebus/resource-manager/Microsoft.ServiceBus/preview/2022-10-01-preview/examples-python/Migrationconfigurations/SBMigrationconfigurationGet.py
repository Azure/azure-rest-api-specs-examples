from azure.identity import DefaultAzureCredential

from azure.mgmt.servicebus import ServiceBusManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-servicebus
# USAGE
    python sb_migrationconfiguration_get.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ServiceBusManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="SubscriptionId",
    )

    response = client.migration_configs.get(
        resource_group_name="ResourceGroup",
        namespace_name="sdk-Namespace-41",
        config_name="$default",
    )
    print(response)


# x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/Migrationconfigurations/SBMigrationconfigurationGet.json
if __name__ == "__main__":
    main()
