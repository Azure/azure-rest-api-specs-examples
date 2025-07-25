from azure.identity import DefaultAzureCredential

from azure.mgmt.policyinsights import PolicyInsightsClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-policyinsights
# USAGE
    python policy_tracked_resources_query_resource_group_scope.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = PolicyInsightsClient(
        credential=DefaultAzureCredential(),
        subscription_id="fffedd8f-ffff-fffd-fffd-fffed2f84852",
    )

    response = client.policy_tracked_resources.list_query_results_for_resource_group(
        resource_group_name="myResourceGroup",
        policy_tracked_resources_resource="default",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/preview/2018-07-01-preview/examples/PolicyTrackedResources_QueryResourceGroupScope.json
if __name__ == "__main__":
    main()
