from azure.identity import DefaultAzureCredential

from azure.mgmt.policyinsights import PolicyInsightsClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-policyinsights
# USAGE
    python policy_restrictions_check_at_subscription_scope.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = PolicyInsightsClient(
        credential=DefaultAzureCredential(),
        subscription_id="35ee058e-5fa0-414c-8145-3ebb8d09b6e2",
    )

    response = client.policy_restrictions.check_at_subscription_scope(
        parameters={
            "pendingFields": [
                {"field": "name", "values": ["myVMName"]},
                {"field": "location", "values": ["eastus", "westus", "westus2", "westeurope"]},
                {"field": "tags"},
            ],
            "resourceDetails": {
                "apiVersion": "2019-12-01",
                "resourceContent": {"properties": {"priority": "Spot"}, "type": "Microsoft.Compute/virtualMachines"},
            },
        },
    )
    print(response)


# x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/PolicyRestrictions_CheckAtSubscriptionScope.json
if __name__ == "__main__":
    main()
