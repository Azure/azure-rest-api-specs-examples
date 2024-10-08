from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.oracledatabase import OracleDatabaseMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-oracledatabase
# USAGE
    python autonomous_database_backup_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = OracleDatabaseMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.autonomous_database_backups.begin_create_or_update(
        resource_group_name="rg000",
        autonomousdatabasename="databasedb1",
        adbbackupid="1711644130",
        resource={
            "properties": {
                "autonomousDatabaseOcid": "ocid1.autonomousdatabase.oc1..aaaaa3klq",
                "displayName": "Nightly Backup",
                "retentionPeriodInDays": 365,
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/autonomousDatabaseBackup_create.json
if __name__ == "__main__":
    main()
