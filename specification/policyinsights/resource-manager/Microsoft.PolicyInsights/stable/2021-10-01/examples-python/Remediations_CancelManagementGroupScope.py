from azure.identity import DefaultAzureCredential
from azure.mgmt.policyinsights import PolicyInsightsClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-policyinsights
# USAGE
    python remediations_cancel_management_group_scope.py

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

    response = client.remediations.cancel_at_management_group(
        management_group_id="financeMg",
        remediation_name="myRemediation",
    )
    print(response)


# x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-10-01/examples/Remediations_CancelManagementGroupScope.json
if __name__ == "__main__":
    main()
