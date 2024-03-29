const { IotCentralClient } = require("@azure/arm-iotcentral");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Check if an IoT Central application subdomain is available.
 *
 * @summary Check if an IoT Central application subdomain is available.
 * x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/Apps_CheckSubdomainAvailability.json
 */
async function appsSubdomainAvailability() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const operationInputs = {
    name: "myiotcentralapp",
    type: "IoTApps",
  };
  const credential = new DefaultAzureCredential();
  const client = new IotCentralClient(credential, subscriptionId);
  const result = await client.apps.checkSubdomainAvailability(operationInputs);
  console.log(result);
}

appsSubdomainAvailability().catch(console.error);
