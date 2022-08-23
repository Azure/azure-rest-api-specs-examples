const { DeviceUpdate } = require("@azure/arm-deviceupdate");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns list of Accounts.
 *
 * @summary Returns list of Accounts.
 * x-ms-original-file: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2022-10-01/examples/Accounts/Accounts_List.json
 */
async function getListOfAccounts() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new DeviceUpdate(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.accounts.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

getListOfAccounts().catch(console.error);
