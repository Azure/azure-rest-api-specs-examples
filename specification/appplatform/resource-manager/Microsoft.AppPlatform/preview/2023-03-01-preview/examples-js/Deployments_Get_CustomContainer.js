const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a Deployment and its properties.
 *
 * @summary Get a Deployment and its properties.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2023-03-01-preview/examples/Deployments_Get_CustomContainer.json
 */
async function deploymentsGetCustomContainer() {
  const subscriptionId =
    process.env["APPPLATFORM_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APPPLATFORM_RESOURCE_GROUP"] || "myResourceGroup";
  const serviceName = "myservice";
  const appName = "myapp";
  const deploymentName = "mydeployment";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.deployments.get(
    resourceGroupName,
    serviceName,
    appName,
    deploymentName
  );
  console.log(result);
}
