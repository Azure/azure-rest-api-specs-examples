from azure.identity import DefaultAzureCredential
from azure.mgmt.hybridnetwork import HybridNetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-hybridnetwork
# USAGE
    python network_function_definition_version_create.py

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

    response = client.network_function_definition_versions.begin_create_or_update(
        resource_group_name="rg",
        publisher_name="TestPublisher",
        network_function_definition_group_name="TestNetworkFunctionDefinitionGroupName",
        network_function_definition_version_name="1.0.0",
        parameters={
            "location": "eastus",
            "properties": {
                "deployParameters": '{"type":"object","properties":{"releaseName":{"type":"string"},"namespace":{"type":"string"}}}',
                "networkFunctionTemplate": {
                    "networkFunctionApplications": [
                        {
                            "artifactProfile": {
                                "artifactStore": {
                                    "id": "/subscriptions/subid/resourcegroups/rg/providers/microsoft.hybridnetwork/publishers/TestPublisher/artifactStores/testArtifactStore"
                                },
                                "helmArtifactProfile": {
                                    "helmPackageName": "fed-rbac",
                                    "helmPackageVersionRange": "~2.1.3",
                                    "imagePullSecretsValuesPaths": ["global.imagePullSecrets"],
                                    "registryValuesPaths": ["global.registry.docker.repoPath"],
                                },
                            },
                            "artifactType": "HelmPackage",
                            "dependsOnProfile": {
                                "installDependsOn": [],
                                "uninstallDependsOn": [],
                                "updateDependsOn": [],
                            },
                            "deployParametersMappingRuleProfile": {
                                "applicationEnablement": "Enabled",
                                "helmMappingRuleProfile": {
                                    "helmPackageVersion": "2.1.3",
                                    "options": {
                                        "installOptions": {"atomic": "true", "timeout": "30", "wait": "waitValue"},
                                        "upgradeOptions": {"atomic": "true", "timeout": "30", "wait": "waitValue"},
                                    },
                                    "releaseName": "{deployParameters.releaseName}",
                                    "releaseNamespace": "{deployParameters.namesapce}",
                                    "values": "",
                                },
                            },
                            "name": "fedrbac",
                        }
                    ],
                    "nfviType": "AzureArcKubernetes",
                },
                "networkFunctionType": "ContainerizedNetworkFunction",
                "versionState": "Active",
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/NetworkFunctionDefinitionVersionCreate.json
if __name__ == "__main__":
    main()
