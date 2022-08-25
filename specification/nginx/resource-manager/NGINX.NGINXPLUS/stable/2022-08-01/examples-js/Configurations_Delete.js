const { NginxManagementClient } = require("@azure/arm-nginx");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Reset the Nginx configuration of given Nginx deployment to default
 *
 * @summary Reset the Nginx configuration of given Nginx deployment to default
 * x-ms-original-file: specification/nginx/resource-manager/NGINX.NGINXPLUS/stable/2022-08-01/examples/Configurations_Delete.json
 */
async function configurationsDelete() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const deploymentName = "myDeployment";
  const configurationName = "default";
  const credential = new DefaultAzureCredential();
  const client = new NginxManagementClient(credential, subscriptionId);
  const result = await client.configurations.beginDeleteAndWait(
    resourceGroupName,
    deploymentName,
    configurationName
  );
  console.log(result);
}

configurationsDelete().catch(console.error);
