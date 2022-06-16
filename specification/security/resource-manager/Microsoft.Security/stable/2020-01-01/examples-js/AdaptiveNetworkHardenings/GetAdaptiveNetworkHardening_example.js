const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a single Adaptive Network Hardening resource
 *
 * @summary Gets a single Adaptive Network Hardening resource
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/AdaptiveNetworkHardenings/GetAdaptiveNetworkHardening_example.json
 */
async function getASingleAdaptiveNetworkHardeningResource() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const resourceGroupName = "rg1";
  const resourceNamespace = "Microsoft.Compute";
  const resourceType = "virtualMachines";
  const resourceName = "vm1";
  const adaptiveNetworkHardeningResourceName = "default";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.adaptiveNetworkHardenings.get(
    resourceGroupName,
    resourceNamespace,
    resourceType,
    resourceName,
    adaptiveNetworkHardeningResourceName
  );
  console.log(result);
}

getASingleAdaptiveNetworkHardeningResource().catch(console.error);
