from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.rdbms.mariadb import MariaDBManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-rdbms
# USAGE
    python server_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = MariaDBManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="ffffffff-ffff-ffff-ffff-ffffffffffff",
    )

    response = client.servers.begin_create(
        resource_group_name="testrg",
        server_name="mariadbtestsvc4",
        parameters={
            "location": "westus",
            "properties": {
                "administratorLogin": "cloudsa",
                "administratorLoginPassword": "<administratorLoginPassword>",
                "createMode": "Default",
                "minimalTlsVersion": "TLS1_2",
                "sslEnforcement": "Enabled",
                "storageProfile": {"backupRetentionDays": 7, "geoRedundantBackup": "Enabled", "storageMB": 128000},
            },
            "sku": {"capacity": 2, "family": "Gen5", "name": "GP_Gen5_2", "tier": "GeneralPurpose"},
            "tags": {"ElasticServer": "1"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/ServerCreate.json
if __name__ == "__main__":
    main()
