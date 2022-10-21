const { ChaosManagementClient } = require("@azure/arm-chaos");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a list of Target resources that extend a tracked regional resource.
 *
 * @summary Get a list of Target resources that extend a tracked regional resource.
 * x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/preview/2022-10-01-preview/examples/ListTargets.json
 */
async function listAllTargetsThatExtendAVirtualMachineResource() {
  const subscriptionId = "6b052e15-03d3-4f17-b2e1-be7f07588291";
  const resourceGroupName = "exampleRG";
  const parentProviderNamespace = "Microsoft.Compute";
  const parentResourceType = "virtualMachines";
  const parentResourceName = "exampleVM";
  const continuationToken = undefined;
  const options = { continuationToken };
  const credential = new DefaultAzureCredential();
  const client = new ChaosManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.targets.list(
    resourceGroupName,
    parentProviderNamespace,
    parentResourceType,
    parentResourceName,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllTargetsThatExtendAVirtualMachineResource().catch(console.error);
