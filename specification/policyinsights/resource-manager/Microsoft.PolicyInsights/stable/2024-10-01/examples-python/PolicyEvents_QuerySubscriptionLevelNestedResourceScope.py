from azure.identity import DefaultAzureCredential

from azure.mgmt.policyinsights import PolicyInsightsClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-policyinsights
# USAGE
    python policy_events_query_subscription_level_nested_resource_scope.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = PolicyInsightsClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.policy_events.list_query_results_for_resource(
        policy_events_resource="default",
        resource_id="subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/providers/Microsoft.SomeNamespace/someResourceType/someResource/someNestedResourceType/someNestedResource",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/PolicyEvents_QuerySubscriptionLevelNestedResourceScope.json
if __name__ == "__main__":
    main()
