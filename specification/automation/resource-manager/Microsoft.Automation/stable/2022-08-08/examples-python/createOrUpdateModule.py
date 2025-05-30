from azure.identity import DefaultAzureCredential

from azure.mgmt.automation import AutomationClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-automation
# USAGE
    python create_or_update_module.py

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

    response = client.module.create_or_update(
        resource_group_name="rg",
        automation_account_name="myAutomationAccount33",
        module_name="OmsCompositeResources",
        parameters={
            "properties": {
                "contentLink": {
                    "contentHash": {
                        "algorithm": "sha265",
                        "value": "07E108A962B81DD9C9BAA89BB47C0F6EE52B29E83758B07795E408D258B2B87A",
                    },
                    "uri": "https://teststorage.blob.core.windows.net/dsccomposite/OmsCompositeResources.zip",
                    "version": "1.0.0.0",
                }
            }
        },
    )
    print(response)


# x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2022-08-08/examples/createOrUpdateModule.json
if __name__ == "__main__":
    main()
