from azure.identity import DefaultAzureCredential
from azure.mgmt.customproviders import Customproviders

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-customproviders
# USAGE
    python list_all_custom_resource_providers_on_the_subscription.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = Customproviders(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.custom_resource_provider.list_by_subscription()
    for item in response:
        print(item)


# x-ms-original-file: specification/customproviders/resource-manager/Microsoft.CustomProviders/preview/2018-09-01-preview/examples/listCustomRPsBySubscription.json
if __name__ == "__main__":
    main()
