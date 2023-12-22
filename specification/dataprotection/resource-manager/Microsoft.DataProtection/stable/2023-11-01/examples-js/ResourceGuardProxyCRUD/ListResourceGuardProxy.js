const { DataProtectionClient } = require("@azure/arm-dataprotection");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the list of ResourceGuardProxies associated with the vault
 *
 * @summary Returns the list of ResourceGuardProxies associated with the vault
 * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-11-01/examples/ResourceGuardProxyCRUD/ListResourceGuardProxy.json
 */
async function getResourceGuardProxies() {
  const subscriptionId =
    process.env["DATAPROTECTION_SUBSCRIPTION_ID"] || "5e13b949-1218-4d18-8b99-7e12155ec4f7";
  const resourceGroupName = process.env["DATAPROTECTION_RESOURCE_GROUP"] || "SampleResourceGroup";
  const vaultName = "sampleVault";
  const credential = new DefaultAzureCredential();
  const client = new DataProtectionClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.dppResourceGuardProxy.list(resourceGroupName, vaultName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
