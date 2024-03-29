from azure.identity import DefaultAzureCredential
from azure.mgmt.subscription import SubscriptionClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-subscription
# USAGE
    python enable_subscription.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SubscriptionClient(
        credential=DefaultAzureCredential(),
    )

    response = client.subscription.enable(
        subscription_id="7948bcee-488c-47ce-941c-38e20ede803d",
    )
    print(response)


# x-ms-original-file: specification/subscription/resource-manager/Microsoft.Subscription/stable/2021-10-01/examples/enableSubscription.json
if __name__ == "__main__":
    main()
