const { PurviewManagementClient } = require("@azure/arm-purview");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of privately linkable resources for an account
 *
 * @summary Gets a list of privately linkable resources for an account
 * x-ms-original-file: specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/PrivateLinkResources_ListByAccount.json
 */
async function privateLinkResourcesListByAccount() {
  const subscriptionId = "12345678-1234-1234-12345678abc";
  const resourceGroupName = "SampleResourceGroup";
  const accountName = "account1";
  const credential = new DefaultAzureCredential();
  const client = new PurviewManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateLinkResources.listByAccount(
    resourceGroupName,
    accountName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

privateLinkResourcesListByAccount().catch(console.error);
