from azure.identity import DefaultAzureCredential
from azure.mgmt.datashare import DataShareManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-datashare
# USAGE
    python provider_share_subscriptions_get_by_share.py

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

    response = client.provider_share_subscriptions.get_by_share(
        resource_group_name="SampleResourceGroup",
        account_name="Account1",
        share_name="Share1",
        provider_share_subscription_id="4256e2cf-0f82-4865-961b-12f83333f487",
    )
    print(response)


# x-ms-original-file: specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/ProviderShareSubscriptions_GetByShare.json
if __name__ == "__main__":
    main()
