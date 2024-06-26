from azure.identity import DefaultAzureCredential
from azure.mgmt.networkanalytics import NetworkAnalyticsMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-networkanalytics
# USAGE
    python data_products_catalogs_list_by_resource_group_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = NetworkAnalyticsMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-00000000000",
    )

    response = client.data_products_catalogs.list_by_resource_group(
        resource_group_name="aoiresourceGroupName",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/DataProductsCatalogs_ListByResourceGroup_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
