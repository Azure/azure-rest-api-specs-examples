from azure.identity import DefaultAzureCredential
from azure.mgmt.sqlvirtualmachine import SqlVirtualMachineManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-sqlvirtualmachine
# USAGE
    python create_or_update_sql_virtual_machine_automated_backup_weekly.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SqlVirtualMachineManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-1111-2222-3333-444444444444",
    )

    response = client.sql_virtual_machines.begin_create_or_update(
        resource_group_name="testrg",
        sql_virtual_machine_name="testvm",
        parameters={
            "location": "northeurope",
            "properties": {
                "autoBackupSettings": {
                    "backupScheduleType": "Manual",
                    "backupSystemDbs": True,
                    "daysOfWeek": ["Monday", "Friday"],
                    "enable": True,
                    "enableEncryption": True,
                    "fullBackupFrequency": "Weekly",
                    "fullBackupStartTime": 6,
                    "fullBackupWindowHours": 11,
                    "logBackupFrequency": 10,
                    "password": "<Password>",
                    "retentionPeriod": 17,
                    "storageAccessKey": "<primary storage access key>",
                    "storageAccountUrl": "https://teststorage.blob.core.windows.net/",
                    "storageContainerName": "testcontainer",
                },
                "autoPatchingSettings": {
                    "dayOfWeek": "Sunday",
                    "enable": True,
                    "maintenanceWindowDuration": 60,
                    "maintenanceWindowStartingHour": 2,
                },
                "keyVaultCredentialSettings": {"enable": False},
                "serverConfigurationsManagementSettings": {
                    "additionalFeaturesServerConfigurations": {"isRServicesEnabled": False},
                    "sqlConnectivityUpdateSettings": {
                        "connectivityType": "PRIVATE",
                        "port": 1433,
                        "sqlAuthUpdatePassword": "<password>",
                        "sqlAuthUpdateUserName": "sqllogin",
                    },
                    "sqlStorageUpdateSettings": {"diskConfigurationType": "NEW", "diskCount": 1, "startingDeviceId": 2},
                    "sqlWorkloadTypeUpdateSettings": {"sqlWorkloadType": "OLTP"},
                },
                "sqlImageSku": "Enterprise",
                "sqlManagement": "Full",
                "sqlServerLicenseType": "PAYG",
                "virtualMachineResourceId": "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachines/testvm",
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-08-01-preview/examples/CreateOrUpdateSqlVirtualMachineAutomatedBackupWeekly.json
if __name__ == "__main__":
    main()
