const { NginxManagementClient } = require("@azure/arm-nginx");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a certificate of given Nginx deployment
 *
 * @summary Get a certificate of given Nginx deployment
 * x-ms-original-file: specification/nginx/resource-manager/NGINX.NGINXPLUS/stable/2022-08-01/examples/Certificates_Get.json
 */
async function certificatesGet() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const deploymentName = "myDeployment";
  const certificateName = "default";
  const credential = new DefaultAzureCredential();
  const client = new NginxManagementClient(credential, subscriptionId);
  const result = await client.certificates.get(resourceGroupName, deploymentName, certificateName);
  console.log(result);
}

certificatesGet().catch(console.error);
