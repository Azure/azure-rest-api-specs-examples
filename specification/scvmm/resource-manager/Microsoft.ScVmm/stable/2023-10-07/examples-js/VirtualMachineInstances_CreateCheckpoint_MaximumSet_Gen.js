const { ScVmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a checkpoint in virtual machine instance.
 *
 * @summary Creates a checkpoint in virtual machine instance.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VirtualMachineInstances_CreateCheckpoint_MaximumSet_Gen.json
 */
async function virtualMachineInstancesCreateCheckpointMaximumSet() {
  const resourceUri = "gtgclehcbsyave";
  const body = {
    name: "ilvltf",
    description: "zoozhfbepldrgpjqsbhpqebtodrhvy",
  };
  const credential = new DefaultAzureCredential();
  const client = new ScVmm(credential);
  const result = await client.virtualMachineInstances.beginCreateCheckpointAndWait(
    resourceUri,
    body,
  );
  console.log(result);
}
