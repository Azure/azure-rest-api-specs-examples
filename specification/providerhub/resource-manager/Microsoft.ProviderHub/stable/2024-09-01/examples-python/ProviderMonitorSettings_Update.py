from azure.identity import DefaultAzureCredential

from azure.mgmt.providerhub import ProviderHubMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-providerhub
# USAGE
    python provider_monitor_settings_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ProviderHubMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.provider_monitor_settings.update(
        resource_group_name="default",
        provider_monitor_setting_name="ContosoMonitorSetting",
    )
    print(response)


# x-ms-original-file: 2024-09-01/ProviderMonitorSettings_Update.json
if __name__ == "__main__":
    main()
