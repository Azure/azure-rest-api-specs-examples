from azure.identity import DefaultAzureCredential

from azure.mgmt.automation import AutomationClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-automation
# USAGE
    python get_hybrid_runbook_worker.py

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

    response = client.hybrid_runbook_workers.get(
        resource_group_name="rg",
        automation_account_name="testaccount",
        hybrid_runbook_worker_group_name="TestHybridGroup",
        hybrid_runbook_worker_id="c010ad12-ef14-4a2a-aa9e-ef22c4745ddd",
    )
    print(response)


# x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2022-08-08/examples/getHybridRunbookWorker.json
if __name__ == "__main__":
    main()
