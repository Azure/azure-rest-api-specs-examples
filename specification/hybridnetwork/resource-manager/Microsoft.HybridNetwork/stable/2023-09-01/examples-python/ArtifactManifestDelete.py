from azure.identity import DefaultAzureCredential
from azure.mgmt.hybridnetwork import HybridNetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-hybridnetwork
# USAGE
    python artifact_manifest_delete.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = HybridNetworkManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    client.artifact_manifests.begin_delete(
        resource_group_name="rg",
        publisher_name="TestPublisher",
        artifact_store_name="TestArtifactStore",
        artifact_manifest_name="TestManifest",
    ).result()


# x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/ArtifactManifestDelete.json
if __name__ == "__main__":
    main()
