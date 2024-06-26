from azure.identity import DefaultAzureCredential
from azure.mgmt.dashboard import DashboardManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-dashboard
# USAGE
    python grafana_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DashboardManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.grafana.begin_create(
        resource_group_name="myResourceGroup",
        workspace_name="myWorkspace",
        request_body_parameters={
            "identity": {"type": "SystemAssigned"},
            "location": "West US",
            "properties": {
                "apiKey": "Enabled",
                "deterministicOutboundIP": "Enabled",
                "enterpriseConfigurations": {"marketplaceAutoRenew": "Enabled", "marketplacePlanId": "myPlanId"},
                "grafanaConfigurations": {
                    "smtp": {
                        "enabled": True,
                        "fromAddress": "test@sendemail.com",
                        "fromName": "emailsender",
                        "host": "smtp.sendemail.com:587",
                        "password": "<password>",
                        "skipVerify": True,
                        "startTLSPolicy": "OpportunisticStartTLS",
                        "user": "username",
                    }
                },
                "grafanaIntegrations": {
                    "azureMonitorWorkspaceIntegrations": [
                        {
                            "azureMonitorWorkspaceResourceId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.monitor/accounts/myAzureMonitorWorkspace"
                        }
                    ]
                },
                "grafanaMajorVersion": "9",
                "grafanaPlugins": {"sample-plugin-id": {}},
                "publicNetworkAccess": "Enabled",
                "zoneRedundancy": "Enabled",
            },
            "sku": {"name": "Standard"},
            "tags": {"Environment": "Dev"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/dashboard/resource-manager/Microsoft.Dashboard/stable/2023-09-01/examples/Grafana_Create.json
if __name__ == "__main__":
    main()
