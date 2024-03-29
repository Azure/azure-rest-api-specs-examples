const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a specific topology component.
 *
 * @summary Gets a specific topology component.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/Topology/GetTopology_example.json
 */
async function getTopology() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "3eeab341-f466-499c-a8be-85427e154bad";
  const resourceGroupName = process.env["SECURITY_RESOURCE_GROUP"] || "myservers";
  const ascLocation = "centralus";
  const topologyResourceName = "vnets";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.topology.get(resourceGroupName, ascLocation, topologyResourceName);
  console.log(result);
}
