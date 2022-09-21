const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Disable remote debugging.
 *
 * @summary Disable remote debugging.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-09-01-preview/examples/Deployments_DisableRemoteDebugging.json
 */
async function deploymentsDisableRemoteDebugging() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const appName = "myapp";
  const deploymentName = "mydeployment";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.deployments.beginDisableRemoteDebuggingAndWait(
    resourceGroupName,
    serviceName,
    appName,
    deploymentName
  );
  console.log(result);
}

deploymentsDisableRemoteDebugging().catch(console.error);
