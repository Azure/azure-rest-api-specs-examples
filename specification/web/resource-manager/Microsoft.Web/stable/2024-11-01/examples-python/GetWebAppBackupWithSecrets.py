from azure.identity import DefaultAzureCredential

from azure.mgmt.web import WebSiteManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-web
# USAGE
    python get_web_app_backup_with_secrets.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = WebSiteManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="34adfa4f-cedf-4dc0-ba29-b6d1a69ab345",
    )

    response = client.web_apps.list_backup_status_secrets(
        resource_group_name="testrg123",
        name="sitef6141",
        backup_id="12345",
        request={
            "properties": {
                "backupName": "abcdwe",
                "backupSchedule": {
                    "frequencyInterval": 7,
                    "frequencyUnit": "Day",
                    "keepAtLeastOneBackup": True,
                    "retentionPeriodInDays": 30,
                    "startTime": "2022-09-02T17:33:11.641Z",
                },
                "databases": [
                    {
                        "connectionString": "DSN=data-source-name[;SERVER=value] [;PWD=value] [;UID=value] [;<Attribute>=<value>]",
                        "connectionStringName": "backend",
                        "databaseType": "SqlAzure",
                        "name": "backenddb",
                    },
                    {
                        "connectionString": "DSN=data-source-name[;SERVER=value] [;PWD=value] [;UID=value] [;<Attribute>=<value>]",
                        "connectionStringName": "stats",
                        "databaseType": "SqlAzure",
                        "name": "statsdb",
                    },
                ],
                "enabled": True,
                "storageAccountUrl": "DefaultEndpointsProtocol=https;AccountName=storagesample;AccountKey=<account-key>",
            }
        },
    )
    print(response)


# x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/GetWebAppBackupWithSecrets.json
if __name__ == "__main__":
    main()
