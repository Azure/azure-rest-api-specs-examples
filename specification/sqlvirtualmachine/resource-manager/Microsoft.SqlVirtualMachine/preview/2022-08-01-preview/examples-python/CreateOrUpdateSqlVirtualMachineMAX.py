from azure.identity import DefaultAzureCredential
from azure.mgmt.sqlvirtualmachine import SqlVirtualMachineManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-sqlvirtualmachine
# USAGE
    python create_or_update_sql_virtual_machine_max.py

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
                "assessmentSettings": {
                    "enable": True,
                    "runImmediately": True,
                    "schedule": {
                        "dayOfWeek": "Sunday",
                        "enable": True,
                        "monthlyOccurrence": None,
                        "startTime": "23:17",
                        "weeklyInterval": 1,
                    },
                },
                "autoBackupSettings": {
                    "backupScheduleType": "Manual",
                    "backupSystemDbs": True,
                    "enable": True,
                    "enableEncryption": True,
                    "fullBackupFrequency": "Daily",
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
                "enableAutomaticUpgrade": True,
                "keyVaultCredentialSettings": {"enable": False},
                "leastPrivilegeMode": "Enabled",
                "serverConfigurationsManagementSettings": {
                    "additionalFeaturesServerConfigurations": {"isRServicesEnabled": False},
                    "azureAdAuthenticationSettings": {"clientId": "11111111-2222-3333-4444-555555555555"},
                    "sqlConnectivityUpdateSettings": {
                        "connectivityType": "PRIVATE",
                        "port": 1433,
                        "sqlAuthUpdatePassword": "<password>",
                        "sqlAuthUpdateUserName": "sqllogin",
                    },
                    "sqlInstanceSettings": {
                        "collation": "SQL_Latin1_General_CP1_CI_AS",
                        "isIfiEnabled": True,
                        "isLpimEnabled": True,
                        "isOptimizeForAdHocWorkloadsEnabled": True,
                        "maxDop": 8,
                        "maxServerMemoryMB": 128,
                        "minServerMemoryMB": 0,
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


# x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-08-01-preview/examples/CreateOrUpdateSqlVirtualMachineMAX.json
if __name__ == "__main__":
    main()
