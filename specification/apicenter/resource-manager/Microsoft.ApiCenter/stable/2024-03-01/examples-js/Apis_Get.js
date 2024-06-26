const { AzureAPICenter } = require("@azure/arm-apicenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns details of the API.
 *
 * @summary Returns details of the API.
 * x-ms-original-file: specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/Apis_Get.json
 */
async function apisGet() {
  const subscriptionId =
    process.env["APICENTER_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APICENTER_RESOURCE_GROUP"] || "contoso-resources";
  const serviceName = "contoso";
  const workspaceName = "default";
  const apiName = "echo-api";
  const credential = new DefaultAzureCredential();
  const client = new AzureAPICenter(credential, subscriptionId);
  const result = await client.apis.get(resourceGroupName, serviceName, workspaceName, apiName);
  console.log(result);
}
