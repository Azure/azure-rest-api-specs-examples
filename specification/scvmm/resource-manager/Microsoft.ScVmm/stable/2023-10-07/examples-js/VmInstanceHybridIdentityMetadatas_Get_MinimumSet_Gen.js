const { ScVmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Implements HybridIdentityMetadata GET method.
 *
 * @summary Implements HybridIdentityMetadata GET method.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VmInstanceHybridIdentityMetadatas_Get_MinimumSet_Gen.json
 */
async function vmInstanceHybridIdentityMetadatasGetMinimumSet() {
  const resourceUri = "gtgclehcbsyave";
  const credential = new DefaultAzureCredential();
  const client = new ScVmm(credential);
  const result = await client.vmInstanceHybridIdentityMetadatas.get(resourceUri);
  console.log(result);
}
