const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Handles requests to list all resources in a subscription.
 *
 * @summary Handles requests to list all resources in a subscription.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/Services_ListBySubscription.json
 */
async function servicesListBySubscription() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.services.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

servicesListBySubscription().catch(console.error);
