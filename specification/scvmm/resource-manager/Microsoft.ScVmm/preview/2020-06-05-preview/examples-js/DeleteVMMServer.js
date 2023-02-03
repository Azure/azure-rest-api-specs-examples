const { Scvmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deboards the SCVMM fabric from Azure.
 *
 * @summary Deboards the SCVMM fabric from Azure.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/DeleteVMMServer.json
 */
async function deleteVmmServer() {
  const subscriptionId =
    process.env["SCVMM_SUBSCRIPTION_ID"] || "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = process.env["SCVMM_RESOURCE_GROUP"] || "testrg";
  const vmmServerName = "ContosoVMMServer";
  const credential = new DefaultAzureCredential();
  const client = new Scvmm(credential, subscriptionId);
  const result = await client.vmmServers.beginDeleteAndWait(resourceGroupName, vmmServerName);
  console.log(result);
}
