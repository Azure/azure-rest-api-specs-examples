from azure.identity import DefaultAzureCredential

from azure.mgmt.automation import AutomationClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-automation
# USAGE
    python get_source_control_sync_job.py

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

    response = client.source_control_sync_job.get(
        resource_group_name="rg",
        automation_account_name="myAutomationAccount33",
        source_control_name="MySourceControl",
        source_control_sync_job_id="ce6fe3e3-9db3-4096-a6b4-82bfb4c10a9a",
    )
    print(response)


# x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2022-08-08/examples/sourceControlSyncJob/getSourceControlSyncJob.json
if __name__ == "__main__":
    main()
