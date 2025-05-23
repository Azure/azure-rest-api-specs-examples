from azure.identity import DefaultAzureCredential

from azure.mgmt.dependencymap import DependencyMapMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-dependencymap
# USAGE
    python discovery_sources_create_or_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DependencyMapMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.discovery_sources.begin_create_or_update(
        resource_group_name="rgdependencyMap",
        map_name="mapsTest1",
        source_name="sourceTest1",
        resource={
            "location": "y",
            "properties": {"sourceId": "wzlrkzumplzjmixbqv", "sourceType": "DiscoverySourceResourceProperties"},
            "tags": {},
        },
    ).result()
    print(response)


# x-ms-original-file: 2025-01-31-preview/DiscoverySources_CreateOrUpdate.json
if __name__ == "__main__":
    main()
