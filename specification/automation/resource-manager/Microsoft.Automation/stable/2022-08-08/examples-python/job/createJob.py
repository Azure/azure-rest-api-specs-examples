from azure.identity import DefaultAzureCredential

from azure.mgmt.automation import AutomationClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-automation
# USAGE
    python create_job.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AutomationClient(
        credential=DefaultAzureCredential(),
        subscription_id="51766542-3ed7-4a72-a187-0c8ab644ddab",
    )

    response = client.job.create(
        resource_group_name="mygroup",
        automation_account_name="ContoseAutomationAccount",
        job_name="foo",
        parameters={
            "properties": {
                "parameters": {"key01": "value01", "key02": "value02"},
                "runOn": "",
                "runbook": {"name": "TestRunbook"},
            }
        },
    )
    print(response)


# x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2022-08-08/examples/job/createJob.json
if __name__ == "__main__":
    main()
