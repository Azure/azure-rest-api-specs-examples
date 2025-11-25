from azure.identity import DefaultAzureCredential

from azure.mgmt.resourcegraph import ResourceGraphClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-resourcegraph
# USAGE
    python graph_query_add.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ResourceGraphClient(
        credential=DefaultAzureCredential(),
    )

    response = client.graph_query.create_or_update(
        subscription_id="024e2271-06fa-46b6-9079-f1ed3c7b070e",
        resource_group_name="my-resource-group",
        resource_name="MyDockerVMs",
        properties={
            "properties": {
                "description": "Docker VMs in PROD",
                "query": "where isnotnull(tags['Prod']) and properties.extensions[0].Name == 'docker'",
            },
            "tags": {},
        },
    )
    print(response)


# x-ms-original-file: specification/resourcegraph/resource-manager/Microsoft.ResourceGraph/stable/2021-03-01/examples/GraphQueryAdd.json
if __name__ == "__main__":
    main()
