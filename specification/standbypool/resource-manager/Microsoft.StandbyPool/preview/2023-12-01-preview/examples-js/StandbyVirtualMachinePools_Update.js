const { StandbyPoolManagementClient } = require("@azure/arm-standbypool");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a StandbyVirtualMachinePoolResource
 *
 * @summary Update a StandbyVirtualMachinePoolResource
 * x-ms-original-file: specification/standbypool/resource-manager/Microsoft.StandbyPool/preview/2023-12-01-preview/examples/StandbyVirtualMachinePools_Update.json
 */
async function standbyVirtualMachinePoolsUpdate() {
  const subscriptionId =
    process.env["STANDBYPOOL_SUBSCRIPTION_ID"] || "8CC31D61-82D7-4B2B-B9DC-6B924DE7D229";
  const resourceGroupName = process.env["STANDBYPOOL_RESOURCE_GROUP"] || "rgstandbypool";
  const standbyVirtualMachinePoolName = "pool";
  const properties = {
    properties: {
      attachedVirtualMachineScaleSetId:
        "/subscriptions/8CC31D61-82D7-4B2B-B9DC-6B924DE7D229/resourceGroups/rgstandbypool/providers/Microsoft.Compute/virtualMachineScaleSets/myVmss",
      elasticityProfile: { maxReadyCapacity: 304 },
      virtualMachineState: "Running",
    },
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new StandbyPoolManagementClient(credential, subscriptionId);
  const result = await client.standbyVirtualMachinePools.update(
    resourceGroupName,
    standbyVirtualMachinePoolName,
    properties,
  );
  console.log(result);
}
