from azure.identity import DefaultAzureCredential
from azure.mgmt.frontdoor import FrontDoorManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-frontdoor
# USAGE
    python waf_list_policies_under_subscription.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = FrontDoorManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.policies.list_by_subscription()
    for item in response:
        print(item)


# x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2022-05-01/examples/WafListPoliciesUnderSubscription.json
if __name__ == "__main__":
    main()
