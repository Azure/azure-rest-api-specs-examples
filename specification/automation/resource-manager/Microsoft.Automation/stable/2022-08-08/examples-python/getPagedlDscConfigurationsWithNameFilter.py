from azure.identity import DefaultAzureCredential

from azure.mgmt.automation import AutomationClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-automation
# USAGE
    python get_pagedl_dsc_configurations_with_name_filter.py

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

    response = client.dsc_configuration.list_by_automation_account(
        resource_group_name="rg",
        automation_account_name="myAutomationAccount33",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2022-08-08/examples/getPagedlDscConfigurationsWithNameFilter.json
if __name__ == "__main__":
    main()
