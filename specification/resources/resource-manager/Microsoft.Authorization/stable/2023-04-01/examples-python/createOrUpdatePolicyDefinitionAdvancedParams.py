from azure.identity import DefaultAzureCredential

from azure.mgmt.resource.policy import PolicyClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-resource
# USAGE
    python create_or_update_policy_definition_advanced_params.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = PolicyClient(
        credential=DefaultAzureCredential(),
        subscription_id="ae640e6b-ba3e-4256-9d62-2993eecfa6f2",
    )

    response = client.policy_definitions.create_or_update(
        policy_definition_name="EventHubDiagnosticLogs",
        parameters={
            "properties": {
                "description": "Audit enabling of logs and retain them up to a year. This enables recreation of activity trails for investigation purposes when a security incident occurs or your network is compromised",
                "displayName": "Event Hubs should have diagnostic logging enabled",
                "metadata": {"category": "Event Hub"},
                "mode": "Indexed",
                "parameters": {
                    "requiredRetentionDays": {
                        "allowedValues": [0, 30, 90, 180, 365],
                        "defaultValue": 365,
                        "metadata": {
                            "description": "The required diagnostic logs retention in days",
                            "displayName": "Required retention (days)",
                        },
                        "type": "Integer",
                    }
                },
                "policyRule": {
                    "if": {"equals": "Microsoft.EventHub/namespaces", "field": "type"},
                    "then": {
                        "details": {
                            "existenceCondition": {
                                "allOf": [
                                    {
                                        "equals": "true",
                                        "field": "Microsoft.Insights/diagnosticSettings/logs[*].retentionPolicy.enabled",
                                    },
                                    {
                                        "equals": "[parameters('requiredRetentionDays')]",
                                        "field": "Microsoft.Insights/diagnosticSettings/logs[*].retentionPolicy.days",
                                    },
                                ]
                            },
                            "type": "Microsoft.Insights/diagnosticSettings",
                        },
                        "effect": "AuditIfNotExists",
                    },
                },
            }
        },
    )
    print(response)


# x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/createOrUpdatePolicyDefinitionAdvancedParams.json
if __name__ == "__main__":
    main()
