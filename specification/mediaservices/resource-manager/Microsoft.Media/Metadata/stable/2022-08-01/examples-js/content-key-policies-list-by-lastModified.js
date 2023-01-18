const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the Content Key Policies in the account
 *
 * @summary Lists the Content Key Policies in the account
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/content-key-policies-list-by-lastModified.json
 */
async function listsContentKeyPoliciesOrderedByLastModified() {
  const subscriptionId =
    process.env["MEDIASERVICES_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["MEDIASERVICES_RESOURCE_GROUP"] || "contoso";
  const accountName = "contosomedia";
  const orderby = "properties/lastModified";
  const options = { orderby };
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.contentKeyPolicies.list(resourceGroupName, accountName, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
