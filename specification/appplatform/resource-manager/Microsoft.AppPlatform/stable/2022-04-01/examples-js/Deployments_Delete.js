const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Operation to delete a Deployment.
 *
 * @summary Operation to delete a Deployment.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/Deployments_Delete.json
 */
async function deploymentsDelete() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const appName = "myapp";
  const deploymentName = "mydeployment";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.deployments.beginDeleteAndWait(
    resourceGroupName,
    serviceName,
    appName,
    deploymentName
  );
  console.log(result);
}

deploymentsDelete().catch(console.error);
