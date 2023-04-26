from azure.identity import DefaultAzureCredential
from azure.mgmt.resourceconnector import ResourceConnectorMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-resourceconnector
# USAGE
    python appliances_list_by_subscription.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ResourceConnectorMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="11111111-2222-3333-4444-555555555555",
    )

    response = client.appliances.list_by_subscription()
    for item in response:
        print(item)


# x-ms-original-file: specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/stable/2022-10-27/examples/AppliancesListBySubscription.json
if __name__ == "__main__":
    main()
