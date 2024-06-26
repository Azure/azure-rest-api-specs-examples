from azure.identity import DefaultAzureCredential
from azure.mgmt.migrationdiscoverysap import MigrationDiscoverySapMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-migrationdiscoverysap
# USAGE
    python sap_discovery_sites_list_by_resource_group.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = MigrationDiscoverySapMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="6d875e77-e412-4d7d-9af4-8895278b4443",
    )

    response = client.sap_discovery_sites.list_by_resource_group(
        resource_group_name="test-rg",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/SAPDiscoverySites/preview/2023-10-01-preview/examples/SAPDiscoverySites_ListByResourceGroup.json
if __name__ == "__main__":
    main()
