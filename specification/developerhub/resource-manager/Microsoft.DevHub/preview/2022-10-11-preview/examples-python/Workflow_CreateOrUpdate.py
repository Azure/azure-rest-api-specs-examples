from azure.identity import DefaultAzureCredential
from azure.mgmt.devhub import DevHubMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-devhub
# USAGE
    python workflow_create_or_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DevHubMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="subscriptionId1",
    )

    response = client.workflow.create_or_update(
        resource_group_name="resourceGroup1",
        workflow_name="workflow1",
        parameters={
            "location": "location1",
            "properties": {
                "githubWorkflowProfile": {
                    "acr": {
                        "acrRegistryName": "registry1",
                        "acrRepositoryName": "repo1",
                        "acrResourceGroup": "resourceGroup1",
                        "acrSubscriptionId": "subscriptionId1",
                    },
                    "aksResourceId": "/subscriptions/subscriptionId1/resourcegroups/resourceGroup1/providers/Microsoft.ContainerService/managedClusters/cluster1",
                    "branchName": "branch1",
                    "deploymentProperties": {
                        "kubeManifestLocations": ["/src/manifests/"],
                        "manifestType": "kube",
                        "overrides": {"key1": "value1"},
                    },
                    "dockerBuildContext": "repo1/src/",
                    "dockerfile": "repo1/images/Dockerfile",
                    "namespace": "namespace1",
                    "oidcCredentials": {
                        "azureClientId": "12345678-3456-7890-5678-012345678901",
                        "azureTenantId": "66666666-3456-7890-5678-012345678901",
                    },
                    "repositoryName": "repo1",
                    "repositoryOwner": "owner1",
                }
            },
            "tags": {"appname": "testApp"},
        },
    )
    print(response)


# x-ms-original-file: specification/developerhub/resource-manager/Microsoft.DevHub/preview/2022-10-11-preview/examples/Workflow_CreateOrUpdate.json
if __name__ == "__main__":
    main()
