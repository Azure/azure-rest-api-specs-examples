const { AzureAPICenter } = require("@azure/arm-apicenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns a collection of API versions.
 *
 * @summary Returns a collection of API versions.
 * x-ms-original-file: specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/ApiVersions_List.json
 */
async function apiVersionsListByApi() {
  const subscriptionId =
    process.env["APICENTER_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APICENTER_RESOURCE_GROUP"] || "contoso-resources";
  const serviceName = "contoso";
  const workspaceName = "default";
  const apiName = "echo-api";
  const credential = new DefaultAzureCredential();
  const client = new AzureAPICenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.apiVersions.list(
    resourceGroupName,
    serviceName,
    workspaceName,
    apiName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
