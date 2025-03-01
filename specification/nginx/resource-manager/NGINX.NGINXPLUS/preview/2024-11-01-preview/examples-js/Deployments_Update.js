const { NginxManagementClient } = require("@azure/arm-nginx");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Update the NGINX deployment
 *
 * @summary Update the NGINX deployment
 * x-ms-original-file: specification/nginx/resource-manager/NGINX.NGINXPLUS/preview/2024-11-01-preview/examples/Deployments_Update.json
 */
async function deploymentsUpdate() {
  const subscriptionId =
    process.env["NGINX_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["NGINX_RESOURCE_GROUP"] || "myResourceGroup";
  const deploymentName = "myDeployment";
  const credential = new DefaultAzureCredential();
  const client = new NginxManagementClient(credential, subscriptionId);
  const result = await client.deployments.beginUpdateAndWait(resourceGroupName, deploymentName);
  console.log(result);
}
