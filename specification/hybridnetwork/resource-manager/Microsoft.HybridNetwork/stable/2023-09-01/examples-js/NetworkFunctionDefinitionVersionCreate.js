const { HybridNetworkManagementClient } = require("@azure/arm-hybridnetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a network function definition version.
 *
 * @summary Creates or updates a network function definition version.
 * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/NetworkFunctionDefinitionVersionCreate.json
 */
async function createOrUpdateANetworkFunctionDefinitionVersionResource() {
  const subscriptionId = process.env["HYBRIDNETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["HYBRIDNETWORK_RESOURCE_GROUP"] || "rg";
  const publisherName = "TestPublisher";
  const networkFunctionDefinitionGroupName = "TestNetworkFunctionDefinitionGroupName";
  const networkFunctionDefinitionVersionName = "1.0.0";
  const parameters = {
    location: "eastus",
    properties: {
      deployParameters:
        '{"type":"object","properties":{"releaseName":{"type":"string"},"namespace":{"type":"string"}}}',
      networkFunctionTemplate: {
        networkFunctionApplications: [
          {
            name: "fedrbac",
            artifactProfile: {
              artifactStore: {
                id: "/subscriptions/subid/resourcegroups/rg/providers/microsoft.hybridnetwork/publishers/TestPublisher/artifactStores/testArtifactStore",
              },
              helmArtifactProfile: {
                helmPackageName: "fed-rbac",
                helmPackageVersionRange: "~2.1.3",
                imagePullSecretsValuesPaths: ["global.imagePullSecrets"],
                registryValuesPaths: ["global.registry.docker.repoPath"],
              },
            },
            artifactType: "HelmPackage",
            dependsOnProfile: {
              installDependsOn: [],
              uninstallDependsOn: [],
              updateDependsOn: [],
            },
            deployParametersMappingRuleProfile: {
              applicationEnablement: "Enabled",
              helmMappingRuleProfile: {
                helmPackageVersion: "2.1.3",
                options: {
                  installOptions: {
                    atomic: "true",
                    timeout: "30",
                    wait: "waitValue",
                  },
                  upgradeOptions: {
                    atomic: "true",
                    timeout: "30",
                    wait: "waitValue",
                  },
                },
                releaseName: "{deployParameters.releaseName}",
                releaseNamespace: "{deployParameters.namesapce}",
                values: "",
              },
            },
          },
        ],
        nfviType: "AzureArcKubernetes",
      },
      networkFunctionType: "ContainerizedNetworkFunction",
      versionState: "Active",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new HybridNetworkManagementClient(credential, subscriptionId);
  const result = await client.networkFunctionDefinitionVersions.beginCreateOrUpdateAndWait(
    resourceGroupName,
    publisherName,
    networkFunctionDefinitionGroupName,
    networkFunctionDefinitionVersionName,
    parameters,
  );
  console.log(result);
}
