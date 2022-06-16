const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Cloud accounts connectors of a subscription
 *
 * @summary Cloud accounts connectors of a subscription
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2020-01-01-preview/examples/Connectors/GetListConnectorSubscription_example.json
 */
async function getAllCloudAccountsConnectorsOfASubscription() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.connectors.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

getAllCloudAccountsConnectorsOfASubscription().catch(console.error);
