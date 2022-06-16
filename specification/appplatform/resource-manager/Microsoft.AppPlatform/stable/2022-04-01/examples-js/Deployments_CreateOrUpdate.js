const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a new Deployment or update an exiting Deployment.
 *
 * @summary Create a new Deployment or update an exiting Deployment.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/Deployments_CreateOrUpdate.json
 */
async function deploymentsCreateOrUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const appName = "myapp";
  const deploymentName = "mydeployment";
  const deploymentResource = {
    properties: {
      deploymentSettings: {
        addonConfigs: {
          applicationConfigurationService: { patterns: ["mypattern"] },
        },
        environmentVariables: { env: "test" },
        resourceRequests: { cpu: "1000m", memory: "3Gi" },
      },
      instances: [],
      source: {
        type: "Source",
        artifactSelector: "sub-module-1",
        relativePath:
          "resources/a172cedcae47474b615c54d510a5d84a8dea3032e958587430b413538be3f333-2019082605-e3095339-1723-44b7-8b5e-31b1003978bc",
        version: "1.0",
      },
    },
    sku: { name: "S0", capacity: 1, tier: "Standard" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.deployments.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serviceName,
    appName,
    deploymentName,
    deploymentResource
  );
  console.log(result);
}

deploymentsCreateOrUpdate().catch(console.error);
