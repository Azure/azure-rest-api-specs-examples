from azure.identity import DefaultAzureCredential
from azure.mgmt.resourcegraph import ResourceGraphClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-resourcegraph
# USAGE
    python resources_facet_query.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ResourceGraphClient(
        credential=DefaultAzureCredential(),
    )

    response = client.resources(
        query={
            "facets": [
                {"expression": "location", "options": {"$top": 3, "sortOrder": "desc"}},
                {"expression": "properties.storageProfile.osDisk.osType", "options": {"$top": 3, "sortOrder": "desc"}},
                {"expression": "nonExistingColumn", "options": {"$top": 3, "sortOrder": "desc"}},
                {
                    "expression": "resourceGroup",
                    "options": {"$top": 3, "sortBy": "tolower(resourceGroup)", "sortOrder": "asc"},
                },
                {"expression": "resourceGroup", "options": {"$top": 3, "filter": "resourceGroup contains 'test'"}},
            ],
            "query": "Resources | where type =~ 'Microsoft.Compute/virtualMachines' | project id, name, location, resourceGroup, properties.storageProfile.osDisk.osType | limit 5",
            "subscriptions": ["cfbbd179-59d2-4052-aa06-9270a38aa9d6"],
        },
    )
    print(response)


# x-ms-original-file: specification/resourcegraph/resource-manager/Microsoft.ResourceGraph/stable/2022-10-01/examples/ResourcesFacetQuery.json
if __name__ == "__main__":
    main()
