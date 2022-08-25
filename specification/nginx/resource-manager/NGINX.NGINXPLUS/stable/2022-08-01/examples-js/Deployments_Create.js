const { NginxManagementClient } = require("@azure/arm-nginx");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update the Nginx deployment
 *
 * @summary Create or update the Nginx deployment
 * x-ms-original-file: specification/nginx/resource-manager/NGINX.NGINXPLUS/stable/2022-08-01/examples/Deployments_Create.json
 */
async function deploymentsCreate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const deploymentName = "myDeployment";
  const credential = new DefaultAzureCredential();
  const client = new NginxManagementClient(credential, subscriptionId);
  const result = await client.deployments.beginCreateAndWait(resourceGroupName, deploymentName);
  console.log(result);
}

deploymentsCreate().catch(console.error);
