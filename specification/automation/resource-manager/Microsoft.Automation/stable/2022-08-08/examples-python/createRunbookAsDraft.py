from azure.identity import DefaultAzureCredential

from azure.mgmt.automation import AutomationClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-automation
# USAGE
    python create_runbook_as_draft.py

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

    response = client.runbook.create_or_update(
        resource_group_name="rg",
        automation_account_name="ContoseAutomationAccount",
        runbook_name="Get-AzureVMTutorial",
        parameters={
            "location": "East US 2",
            "name": "Get-AzureVMTutorial",
            "properties": {
                "description": "Description of the Runbook",
                "draft": {},
                "logProgress": False,
                "logVerbose": False,
                "runbookType": "PowerShellWorkflow",
            },
            "tags": {"tag01": "value01", "tag02": "value02"},
        },
    )
    print(response)


# x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2022-08-08/examples/createRunbookAsDraft.json
if __name__ == "__main__":
    main()
