const { ScVmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to create or update a virtual machine instance. Please note some properties can be set only during virtual machine instance creation.
 *
 * @summary The operation to create or update a virtual machine instance. Please note some properties can be set only during virtual machine instance creation.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VirtualMachineInstances_CreateOrUpdate_MinimumSet_Gen.json
 */
async function virtualMachineInstancesCreateOrUpdateMinimumSet() {
  const resourceUri = "gtgclehcbsyave";
  const resource = { extendedLocation: {} };
  const credential = new DefaultAzureCredential();
  const client = new ScVmm(credential);
  const result = await client.virtualMachineInstances.beginCreateOrUpdateAndWait(
    resourceUri,
    resource,
  );
  console.log(result);
}
