from azure.identity import DefaultAzureCredential
from azure.mgmt.resourcegraph import ResourceGraphClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-resourcegraph
# USAGE
    python resources_history_mgs_get.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ResourceGraphClient(
        credential=DefaultAzureCredential(),
    )

    response = client.resources_history(
        request={
            "managementGroups": ["e927f598-c1d4-4f72-8541-95d83a6a4ac8", "ProductionMG"],
            "options": {"interval": {"end": "2020-11-12T01:25:00.0000000Z", "start": "2020-11-12T01:00:00.0000000Z"}},
            "query": "where name =~ 'cpu-utilization' | project id, name, properties",
        },
    )
    print(response)


# x-ms-original-file: specification/resourcegraph/resource-manager/Microsoft.ResourceGraph/preview/2021-06-01-preview/examples/ResourcesHistoryMgsGet.json
if __name__ == "__main__":
    main()
