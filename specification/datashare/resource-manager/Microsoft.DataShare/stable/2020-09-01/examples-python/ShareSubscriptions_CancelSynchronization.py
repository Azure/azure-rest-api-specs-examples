from azure.identity import DefaultAzureCredential
from azure.mgmt.datashare import DataShareManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-datashare
# USAGE
    python share_subscriptions_cancel_synchronization.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DataShareManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="12345678-1234-1234-12345678abc",
    )

    response = client.share_subscriptions.begin_cancel_synchronization(
        resource_group_name="SampleResourceGroup",
        account_name="Account1",
        share_subscription_name="ShareSubscription1",
        share_subscription_synchronization={"synchronizationId": "7d0536a6-3fa5-43de-b152-3d07c4f6b2bb"},
    ).result()
    print(response)


# x-ms-original-file: specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/ShareSubscriptions_CancelSynchronization.json
if __name__ == "__main__":
    main()
