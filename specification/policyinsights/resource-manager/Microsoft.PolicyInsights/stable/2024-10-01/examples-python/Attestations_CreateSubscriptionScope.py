from azure.identity import DefaultAzureCredential

from azure.mgmt.policyinsights import PolicyInsightsClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-policyinsights
# USAGE
    python attestations_create_subscription_scope.py

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

    response = client.attestations.begin_create_or_update_at_subscription(
        attestation_name="790996e6-9871-4b1f-9cd9-ec42cd6ced1e",
        parameters={
            "properties": {
                "complianceState": "Compliant",
                "policyAssignmentId": "/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/providers/microsoft.authorization/policyassignments/b101830944f246d8a14088c5",
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/Attestations_CreateSubscriptionScope.json
if __name__ == "__main__":
    main()
