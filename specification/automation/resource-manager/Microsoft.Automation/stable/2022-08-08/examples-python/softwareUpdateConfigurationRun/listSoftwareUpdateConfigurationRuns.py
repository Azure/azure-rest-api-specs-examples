from azure.identity import DefaultAzureCredential

from azure.mgmt.automation import AutomationClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-automation
# USAGE
    python list_software_update_configuration_runs.py

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

    response = client.software_update_configuration_runs.list(
        resource_group_name="mygroup",
        automation_account_name="myaccount",
    )
    print(response)


# x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2022-08-08/examples/softwareUpdateConfigurationRun/listSoftwareUpdateConfigurationRuns.json
if __name__ == "__main__":
    main()
