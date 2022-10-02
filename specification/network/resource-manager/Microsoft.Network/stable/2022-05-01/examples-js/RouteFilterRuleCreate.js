const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a route in the specified route filter.
 *
 * @summary Creates or updates a route in the specified route filter.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/RouteFilterRuleCreate.json
 */
async function routeFilterRuleCreate() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const routeFilterName = "filterName";
  const ruleName = "ruleName";
  const routeFilterRuleParameters = {
    access: "Allow",
    communities: ["12076:5030", "12076:5040"],
    routeFilterRuleType: "Community",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.routeFilterRules.beginCreateOrUpdateAndWait(
    resourceGroupName,
    routeFilterName,
    ruleName,
    routeFilterRuleParameters
  );
  console.log(result);
}

routeFilterRuleCreate().catch(console.error);
