from azure.identity import DefaultAzureCredential

from azure.mgmt.azurestackhci import AzureStackHCIClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-azurestackhci
# USAGE
    python put_update_summaries.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AzureStackHCIClient(
        credential=DefaultAzureCredential(),
        subscription_id="b8d594e5-51f3-4c11-9c54-a7771b81c712",
    )

    response = client.update_summaries.put(
        resource_group_name="testrg",
        cluster_name="testcluster",
        update_location_properties={
            "properties": {
                "currentVersion": "4.2203.2.32",
                "hardwareModel": "PowerEdge R730xd",
                "lastChecked": "2022-04-07T18:04:07Z",
                "lastUpdated": "2022-04-06T14:08:18.254Z",
                "oemFamily": "DellEMC",
                "state": "AppliedSuccessfully",
            }
        },
    )
    print(response)


# x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/PutUpdateSummaries.json
if __name__ == "__main__":
    main()
