const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Enable remote debugging.
 *
 * @summary Enable remote debugging.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/Deployments_EnableRemoteDebugging.json
 */
async function deploymentsEnableRemoteDebugging() {
  const subscriptionId =
    process.env["APPPLATFORM_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APPPLATFORM_RESOURCE_GROUP"] || "myResourceGroup";
  const serviceName = "myservice";
  const appName = "myapp";
  const deploymentName = "mydeployment";
  const remoteDebuggingPayload = {};
  const options = {
    remoteDebuggingPayload,
  };
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.deployments.beginEnableRemoteDebuggingAndWait(
    resourceGroupName,
    serviceName,
    appName,
    deploymentName,
    options,
  );
  console.log(result);
}
