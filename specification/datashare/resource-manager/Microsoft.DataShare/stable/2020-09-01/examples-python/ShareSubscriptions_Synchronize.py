from azure.identity import DefaultAzureCredential
from azure.mgmt.datashare import DataShareManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-datashare
# USAGE
    python share_subscriptions_synchronize.py

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

    response = client.share_subscriptions.begin_synchronize(
        resource_group_name="SampleResourceGroup",
        account_name="Account1",
        share_subscription_name="ShareSubscription1",
        synchronize={"synchronizationMode": "Incremental"},
    ).result()
    print(response)


# x-ms-original-file: specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/ShareSubscriptions_Synchronize.json
if __name__ == "__main__":
    main()
