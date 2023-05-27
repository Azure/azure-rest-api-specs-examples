from azure.identity import DefaultAzureCredential
from azure.mgmt.costmanagement import CostManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-costmanagement
# USAGE
    python external_subscriptions_dimensions.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = CostManagementClient(
        credential=DefaultAzureCredential(),
    )

    response = client.dimensions.by_external_cloud_provider_type(
        external_cloud_provider_type="externalSubscriptions",
        external_cloud_provider_id="100",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ExternalSubscriptionsDimensions.json
if __name__ == "__main__":
    main()
