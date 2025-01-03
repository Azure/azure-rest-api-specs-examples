from azure.identity import DefaultAzureCredential

from azure.mgmt.automation import AutomationClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-automation
# USAGE
    python create_job_schedule.py

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

    response = client.job_schedule.create(
        resource_group_name="rg",
        automation_account_name="ContoseAutomationAccount",
        job_schedule_id="0fa462ba-3aa2-4138-83ca-9ebc3bc55cdc",
        parameters={
            "properties": {
                "parameters": {"jobscheduletag01": "jobschedulevalue01", "jobscheduletag02": "jobschedulevalue02"},
                "runbook": {"name": "TestRunbook"},
                "schedule": {"name": "ScheduleNameGoesHere332204b5-debe-4348-a5c7-6357457189f2"},
            }
        },
    )
    print(response)


# x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2022-08-08/examples/createJobSchedule.json
if __name__ == "__main__":
    main()
