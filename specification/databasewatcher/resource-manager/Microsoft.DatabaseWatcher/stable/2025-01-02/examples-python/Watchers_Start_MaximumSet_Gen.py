from azure.identity import DefaultAzureCredential

from azure.mgmt.databasewatcher import DatabaseWatcherMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-databasewatcher
# USAGE
    python watchers_start_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DatabaseWatcherMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    client.watchers.begin_start(
        resource_group_name="rgWatcher",
        watcher_name="testWatcher",
    ).result()


# x-ms-original-file: 2025-01-02/Watchers_Start_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
