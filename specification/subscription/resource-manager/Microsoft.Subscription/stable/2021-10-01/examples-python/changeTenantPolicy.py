from azure.identity import DefaultAzureCredential
from azure.mgmt.subscription import SubscriptionClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-subscription
# USAGE
    python change_tenant_policy.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SubscriptionClient(
        credential=DefaultAzureCredential(),
    )

    response = client.subscription_policy.add_update_policy_for_tenant(
        body={
            "blockSubscriptionsIntoTenant": True,
            "blockSubscriptionsLeavingTenant": True,
            "exemptedPrincipals": ["e879cf0f-2b4d-5431-109a-f72fc9868693", "9792da87-c97b-410d-a97d-27021ba09ce6"],
        },
    )
    print(response)


# x-ms-original-file: specification/subscription/resource-manager/Microsoft.Subscription/stable/2021-10-01/examples/changeTenantPolicy.json
if __name__ == "__main__":
    main()
