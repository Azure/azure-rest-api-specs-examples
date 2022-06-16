const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Enforces the given rules on the NSG(s) listed in the request
 *
 * @summary Enforces the given rules on the NSG(s) listed in the request
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/AdaptiveNetworkHardenings/EnforceAdaptiveNetworkHardeningRules_example.json
 */
async function enforcesTheGivenRulesOnTheNsgSListedInTheRequest() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const resourceGroupName = "rg1";
  const resourceNamespace = "Microsoft.Compute";
  const resourceType = "virtualMachines";
  const resourceName = "vm1";
  const adaptiveNetworkHardeningResourceName = "default";
  const body = {
    networkSecurityGroups: [
      "/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/nsg1",
      "/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/rg2/providers/Microsoft.Network/networkSecurityGroups/nsg2",
    ],
    rules: [
      {
        name: "rule1",
        destinationPort: 3389,
        direction: "Inbound",
        ipAddresses: ["100.10.1.1", "200.20.2.2", "81.199.3.0/24"],
        protocols: ["TCP"],
      },
      {
        name: "rule2",
        destinationPort: 22,
        direction: "Inbound",
        ipAddresses: [],
        protocols: ["TCP"],
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.adaptiveNetworkHardenings.beginEnforceAndWait(
    resourceGroupName,
    resourceNamespace,
    resourceType,
    resourceName,
    adaptiveNetworkHardeningResourceName,
    body
  );
  console.log(result);
}

enforcesTheGivenRulesOnTheNsgSListedInTheRequest().catch(console.error);
