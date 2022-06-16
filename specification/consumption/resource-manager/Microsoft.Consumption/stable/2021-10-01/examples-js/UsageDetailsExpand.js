const { ConsumptionManagementClient } = require("@azure/arm-consumption");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the usage details for the defined scope. Usage details are available via this API only for May 1, 2014 or later.
 *
 * @summary Lists the usage details for the defined scope. Usage details are available via this API only for May 1, 2014 or later.
 * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/UsageDetailsExpand.json
 */
async function usageDetailsExpandLegacy() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const scope = "subscriptions/00000000-0000-0000-0000-000000000000";
  const expand = "meterDetails,additionalInfo";
  const filter = "tags eq 'dev:tools'";
  const top = 1;
  const options = { expand, filter, top };
  const credential = new DefaultAzureCredential();
  const client = new ConsumptionManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.usageDetails.list(scope, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

usageDetailsExpandLegacy().catch(console.error);
