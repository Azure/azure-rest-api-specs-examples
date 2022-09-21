const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List deployments for a certain service
 *
 * @summary List deployments for a certain service
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-09-01-preview/examples/Deployments_ListForCluster.json
 */
async function deploymentsListForCluster() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.deployments.listForCluster(resourceGroupName, serviceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

deploymentsListForCluster().catch(console.error);
