from azure.identity import DefaultAzureCredential

from azure.mgmt.recoveryservicessiterecovery import SiteRecoveryManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-recoveryservicessiterecovery
# USAGE
    python replication_alert_settings_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SiteRecoveryManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.replication_alert_settings.create(
        resource_group_name="resourceGroupPS1",
        resource_name="vault1",
        alert_setting_name="defaultAlertSetting",
        request={
            "properties": {"customEmailAddresses": ["ronehr@microsoft.com"], "locale": "", "sendToOwners": "false"}
        },
    )
    print(response)


# x-ms-original-file: 2025-08-01/ReplicationAlertSettings_Create.json
if __name__ == "__main__":
    main()
