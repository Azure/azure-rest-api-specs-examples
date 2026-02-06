from azure.identity import DefaultAzureCredential

from azure.mgmt.resource.subscriptions import SubscriptionClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-resource-subscriptions
# USAGE
    python get_tenants.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SubscriptionClient(
        credential=DefaultAzureCredential(),
    )

    response = client.tenants.list()
    for item in response:
        print(item)


# x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/subscriptions/stable/2022-12-01/examples/GetTenants.json
if __name__ == "__main__":
    main()
