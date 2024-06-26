from azure.identity import DefaultAzureCredential
from azure.mgmt.subscription import SubscriptionClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-subscription
# USAGE
    python cancel_subscription.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SubscriptionClient(
        credential=DefaultAzureCredential(),
    )

    response = client.subscription.cancel(
        subscription_id="83aa47df-e3e9-49ff-877b-94304bf3d3ad",
    )
    print(response)


# x-ms-original-file: specification/subscription/resource-manager/Microsoft.Subscription/stable/2021-10-01/examples/cancelSubscription.json
if __name__ == "__main__":
    main()
