from azure.identity import DefaultAzureCredential

from azure.mgmt.mysqlflexibleservers import MySQLManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-mysqlflexibleservers
# USAGE
    python advanced_threat_protection_settings_list.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = MySQLManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-1111-2222-3333-444444444444",
    )

    response = client.advanced_threat_protection_settings.list(
        resource_group_name="threatprotection-6852",
        server_name="threatprotection-2080",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/FlexibleServers/preview/2024-10-01-preview/examples/AdvancedThreatProtectionSettingsList.json
if __name__ == "__main__":
    main()
