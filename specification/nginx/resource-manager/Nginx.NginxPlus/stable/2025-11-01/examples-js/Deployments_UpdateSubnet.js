const { NginxManagementClient } = require("@azure/arm-nginx");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update the NGINX deployment
 *
 * @summary update the NGINX deployment
 * x-ms-original-file: 2025-11-01/Deployments_UpdateSubnet.json
 */
async function deploymentsUpdateSubnet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NginxManagementClient(credential, subscriptionId);
  const result = await client.deployments.update("myResourceGroup", "myDeployment");
  console.log(result);
}
