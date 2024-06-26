from azure.identity import DefaultAzureCredential
from azure.mgmt.migrationdiscoverysap import MigrationDiscoverySapMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-migrationdiscoverysap
# USAGE
    python sap_instances_delete.py

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

    client.sap_instances.begin_delete(
        resource_group_name="test-rg",
        sap_discovery_site_name="SampleSite",
        sap_instance_name="MPP_MPP",
    ).result()


# x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/SAPDiscoverySites/preview/2023-10-01-preview/examples/SAPInstances_Delete.json
if __name__ == "__main__":
    main()
