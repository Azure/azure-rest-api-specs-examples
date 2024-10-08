from azure.identity import DefaultAzureCredential

from azure.mgmt.mysqlflexibleservers import MySQLManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-mysqlflexibleservers
# USAGE
    python configurations_batch_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = MySQLManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="ffffffff-ffff-ffff-ffff-ffffffffffff",
    )

    response = client.configurations.begin_batch_update(
        resource_group_name="testrg",
        server_name="mysqltestserver",
        parameters={
            "resetAllToDefault": "False",
            "value": [
                {"name": "event_scheduler", "properties": {"value": "OFF"}},
                {"name": "div_precision_increment", "properties": {"value": "8"}},
            ],
        },
    ).result()
    print(response)


# x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/Configurations/stable/2023-12-30/examples/ConfigurationsBatchUpdate.json
if __name__ == "__main__":
    main()
