from azure.identity import DefaultAzureCredential

from azure.mgmt.automation import AutomationClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-automation
# USAGE
    python private_endpoint_connection_get.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AutomationClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.private_endpoint_connections.get(
        resource_group_name="rg1",
        automation_account_name="automationAccountName",
        private_endpoint_connection_name="privateEndpointConnectionName",
    )
    print(response)


# x-ms-original-file: 2024-10-23/privateEndpointConnection/PrivateEndpointConnectionGet.json
if __name__ == "__main__":
    main()
