const { NginxManagementClient } = require("@azure/arm-nginx");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all operations provided by Nginx.NginxPlus for the 2023-04-01 api version.
 *
 * @summary List all operations provided by Nginx.NginxPlus for the 2023-04-01 api version.
 * x-ms-original-file: specification/nginx/resource-manager/NGINX.NGINXPLUS/stable/2023-04-01/examples/Operations_List.json
 */
async function operationsList() {
  const subscriptionId =
    process.env["NGINX_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new NginxManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.operations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
