from azure.identity import DefaultAzureCredential

from azure.mgmt.cloudhealth import CloudHealthMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-cloudhealth
# USAGE
    python entities_create_or_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = CloudHealthMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.entities.create_or_update(
        resource_group_name="rgopenapi",
        health_model_name="myHealthModel",
        entity_name="uszrxbdkxesdrxhmagmzywebgbjj",
        resource={
            "properties": {
                "alerts": {
                    "degraded": {
                        "actionGroupIds": [
                            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Insights/actionGroups/myactiongroup"
                        ],
                        "description": "Alert description",
                        "severity": "Sev4",
                    },
                    "unhealthy": {
                        "actionGroupIds": [
                            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Insights/actionGroups/myactiongroup"
                        ],
                        "description": "Alert description",
                        "severity": "Sev1",
                    },
                },
                "canvasPosition": {"x": 14, "y": 13},
                "displayName": "My entity",
                "healthObjective": 62,
                "icon": {"customData": "rcitntvapruccrhtxmkqjphbxunkz", "iconName": "Custom"},
                "impact": "Standard",
                "labels": {"key1376": "ixfvzsfnpvkkbrce"},
                "signals": {
                    "azureLogAnalytics": {
                        "authenticationSetting": "B3P1X3e-FZtZ-4Ak-2VLHGQ-4m4-05DE-XNW5zW3P-46XY-DC3SSX",
                        "logAnalyticsWorkspaceResourceId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.OperationalInsights/workspaces/myworkspace",
                        "signalAssignments": [
                            {"signalDefinitions": ["B3P1X3e-FZtZ-4Ak-2VLHGQ-4m4-05DE-XNW5zW3P-46XY-DC3SSX"]}
                        ],
                    },
                    "azureMonitorWorkspace": {
                        "authenticationSetting": "B3P1X3e-FZtZ-4Ak-2VLHGQ-4m4-05DE-XNW5zW3P-46XY-DC3SSX",
                        "azureMonitorWorkspaceResourceId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.OperationalInsights/workspaces/myworkspace",
                        "signalAssignments": [{"signalDefinitions": ["sigdef2"]}, {"signalDefinitions": ["sigdef3"]}],
                    },
                    "azureResource": {
                        "authenticationSetting": "B3P1X3e-FZtZ-4Ak-2VLHGQ-4m4-05DE-XNW5zW3P-46XY-DC3SSX",
                        "azureResourceId": "/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/vm1",
                        "signalAssignments": [{"signalDefinitions": ["sigdef1"]}],
                    },
                    "dependencies": {"aggregationType": "WorstOf"},
                },
            }
        },
    )
    print(response)


# x-ms-original-file: 2025-05-01-preview/Entities_CreateOrUpdate.json
if __name__ == "__main__":
    main()
