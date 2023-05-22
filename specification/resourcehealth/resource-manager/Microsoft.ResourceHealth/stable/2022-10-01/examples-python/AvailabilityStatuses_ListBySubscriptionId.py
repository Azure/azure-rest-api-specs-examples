from azure.identity import DefaultAzureCredential
from azure.mgmt.resourcehealth import ResourceHealthMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-resourcehealth
# USAGE
    python availability_statuses_list_by_subscription_id.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ResourceHealthMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="subscriptionId",
    )

    response = client.availability_statuses.list_by_subscription_id()
    for item in response:
        print(item)


# x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/stable/2022-10-01/examples/AvailabilityStatuses_ListBySubscriptionId.json
if __name__ == "__main__":
    main()
