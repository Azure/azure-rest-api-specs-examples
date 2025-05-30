from azure.identity import DefaultAzureCredential

from azure.mgmt.automation import AutomationClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-automation
# USAGE
    python create_or_update_watcher.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AutomationClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.watcher.create_or_update(
        resource_group_name="rg",
        automation_account_name="MyTestAutomationAccount",
        watcher_name="MyTestWatcher",
        parameters={
            "etag": None,
            "location": None,
            "properties": {
                "description": "This is a test watcher.",
                "executionFrequencyInSeconds": 60,
                "lastModifiedBy": None,
                "scriptName": "MyTestWatcherRunbook",
                "scriptParameters": None,
                "scriptRunOn": "MyTestHybridWorkerGroup",
            },
            "tags": {},
            "type": None,
        },
    )
    print(response)


# x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/createOrUpdateWatcher.json
if __name__ == "__main__":
    main()
