from azure.identity import DefaultAzureCredential

from azure.mgmt.automation import AutomationClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-automation
# USAGE
    python create_or_update_connection.py

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

    response = client.connection.create_or_update(
        resource_group_name="rg",
        automation_account_name="myAutomationAccount28",
        connection_name="mysConnection",
        parameters={
            "name": "mysConnection",
            "properties": {
                "connectionType": {"name": "Azure"},
                "description": "my description goes here",
                "fieldDefinitionValues": {"AutomationCertificateName": "mysCertificateName", "SubscriptionID": "subid"},
            },
        },
    )
    print(response)


# x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2022-08-08/examples/createOrUpdateConnection.json
if __name__ == "__main__":
    main()
