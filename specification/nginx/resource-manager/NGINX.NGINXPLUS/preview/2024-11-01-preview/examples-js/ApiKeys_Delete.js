const { NginxManagementClient } = require("@azure/arm-nginx");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Delete API key for Nginx deployment
 *
 * @summary Delete API key for Nginx deployment
 * x-ms-original-file: specification/nginx/resource-manager/NGINX.NGINXPLUS/preview/2024-11-01-preview/examples/ApiKeys_Delete.json
 */
async function apiKeysDelete() {
  const subscriptionId =
    process.env["NGINX_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["NGINX_RESOURCE_GROUP"] || "myResourceGroup";
  const deploymentName = "myDeployment";
  const apiKeyName = "myApiKey";
  const credential = new DefaultAzureCredential();
  const client = new NginxManagementClient(credential, subscriptionId);
  const result = await client.apiKeys.delete(resourceGroupName, deploymentName, apiKeyName);
  console.log(result);
}
