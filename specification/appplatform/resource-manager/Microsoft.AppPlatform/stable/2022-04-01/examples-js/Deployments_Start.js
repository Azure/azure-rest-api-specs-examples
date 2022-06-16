const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Start the deployment.
 *
 * @summary Start the deployment.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/Deployments_Start.json
 */
async function deploymentsStart() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const appName = "myapp";
  const deploymentName = "mydeployment";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.deployments.beginStartAndWait(
    resourceGroupName,
    serviceName,
    appName,
    deploymentName
  );
  console.log(result);
}

deploymentsStart().catch(console.error);
