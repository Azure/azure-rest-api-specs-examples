from azure.identity import DefaultAzureCredential

from azure.mgmt.loganalytics import LogAnalyticsManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-loganalytics
# USAGE
    python query_packs_get.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = LogAnalyticsManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="86dc51d3-92ed-4d7e-947a-775ea79b4919",
    )

    response = client.query_packs.get(
        resource_group_name="my-resource-group",
        query_pack_name="my-querypack",
    )
    print(response)


# x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2019-09-01/examples/QueryPacksGet.json
if __name__ == "__main__":
    main()
