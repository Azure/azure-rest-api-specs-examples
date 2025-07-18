from azure.identity import DefaultAzureCredential

from azure.mgmt.servicefabricmanagedclusters import ServiceFabricManagedClustersManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-servicefabricmanagedclusters
# USAGE
    python managed_cluster_version_get_by_environment_example.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ServiceFabricManagedClustersManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.managed_cluster_version.get_by_environment(
        location="eastus",
        environment="Windows",
        cluster_version="7.2.477.9590",
    )
    print(response)


# x-ms-original-file: 2025-03-01-preview/ManagedClusterVersionGetByEnvironment_example.json
if __name__ == "__main__":
    main()
