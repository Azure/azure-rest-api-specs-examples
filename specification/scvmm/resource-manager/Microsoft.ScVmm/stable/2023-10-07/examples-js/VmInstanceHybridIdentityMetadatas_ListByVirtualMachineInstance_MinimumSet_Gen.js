const { ScVmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the list of HybridIdentityMetadata of the given VM.
 *
 * @summary Returns the list of HybridIdentityMetadata of the given VM.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VmInstanceHybridIdentityMetadatas_ListByVirtualMachineInstance_MinimumSet_Gen.json
 */
async function vmInstanceHybridIdentityMetadatasListByVirtualMachineInstanceMinimumSet() {
  const resourceUri = "gtgclehcbsyave";
  const credential = new DefaultAzureCredential();
  const client = new ScVmm(credential);
  const resArray = new Array();
  for await (let item of client.vmInstanceHybridIdentityMetadatas.listByVirtualMachineInstance(
    resourceUri,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
