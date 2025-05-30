from azure.identity import DefaultAzureCredential

from azure.mgmt.resource.policy import PolicyClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-resource
# USAGE
    python create_or_update_policy_definition.py

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
        policy_definition_name="ResourceNaming",
        parameters={
            "properties": {
                "description": "Force resource names to begin with given 'prefix' and/or end with given 'suffix'",
                "displayName": "Enforce resource naming convention",
                "metadata": {"category": "Naming"},
                "mode": "All",
                "parameters": {
                    "prefix": {
                        "metadata": {"description": "Resource name prefix", "displayName": "Prefix"},
                        "type": "String",
                    },
                    "suffix": {
                        "metadata": {"description": "Resource name suffix", "displayName": "Suffix"},
                        "type": "String",
                    },
                },
                "policyRule": {
                    "if": {
                        "not": {"field": "name", "like": "[concat(parameters('prefix'), '*', parameters('suffix'))]"}
                    },
                    "then": {"effect": "deny"},
                },
            }
        },
    )
    print(response)


# x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/createOrUpdatePolicyDefinition.json
if __name__ == "__main__":
    main()
