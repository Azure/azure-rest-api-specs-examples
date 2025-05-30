from azure.identity import DefaultAzureCredential

from azure.mgmt.automation import AutomationClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-automation
# USAGE
    python private_endpoint_connection_delete.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AutomationClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-1111-2222-3333-444444444444",
    )

    client.private_endpoint_connections.begin_delete(
        resource_group_name="rg1",
        automation_account_name="ddb1",
        private_endpoint_connection_name="privateEndpointConnectionName",
    ).result()


# x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/PrivateEndpointConnectionDelete.json
if __name__ == "__main__":
    main()
