from azure.identity import DefaultAzureCredential

from azure.mgmt.sql import SqlManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-sql
# USAGE
    python database_advanced_threat_protection_settings_create_min.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SqlManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-1111-2222-3333-444444444444",
    )

    response = client.database_advanced_threat_protection_settings.create_or_update(
        resource_group_name="threatprotection-4799",
        server_name="threatprotection-6440",
        database_name="testdb",
        advanced_threat_protection_name="Default",
        parameters={"properties": {"state": "Disabled"}},
    )
    print(response)


# x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2024-05-01-preview/examples/DatabaseAdvancedThreatProtectionSettingsCreateMin.json
if __name__ == "__main__":
    main()
