from azure.identity import DefaultAzureCredential
from azure.mgmt.hybridnetwork import HybridNetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-hybridnetwork
# USAGE
    python artifact_manifest_create.py

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

    response = client.artifact_manifests.begin_create_or_update(
        resource_group_name="rg",
        publisher_name="TestPublisher",
        artifact_store_name="TestArtifactStore",
        artifact_manifest_name="TestManifest",
        parameters={
            "location": "eastus",
            "properties": {
                "artifacts": [
                    {"artifactName": "fed-rbac", "artifactType": "OCIArtifact", "artifactVersion": "1.0.0"},
                    {"artifactName": "nginx", "artifactType": "OCIArtifact", "artifactVersion": "v1"},
                ]
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/ArtifactManifestCreate.json
if __name__ == "__main__":
    main()
