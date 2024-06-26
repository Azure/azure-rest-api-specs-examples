const { ScVmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Removes the SCVmm fabric from Azure.
 *
 * @summary Removes the SCVmm fabric from Azure.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VmmServers_Delete_MaximumSet_Gen.json
 */
async function vmmServersDeleteMaximumSet() {
  const subscriptionId =
    process.env["SCVMM_SUBSCRIPTION_ID"] || "79332E5A-630B-480F-A266-A941C015AB19";
  const resourceGroupName = process.env["SCVMM_RESOURCE_GROUP"] || "rgscvmm";
  const force = "true";
  const vmmServerName = ".";
  const options = { force };
  const credential = new DefaultAzureCredential();
  const client = new ScVmm(credential, subscriptionId);
  const result = await client.vmmServers.beginDeleteAndWait(
    resourceGroupName,
    vmmServerName,
    options,
  );
  console.log(result);
}
