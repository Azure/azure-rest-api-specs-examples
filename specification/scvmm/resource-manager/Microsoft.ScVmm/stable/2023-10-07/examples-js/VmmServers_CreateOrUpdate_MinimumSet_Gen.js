const { ScVmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Onboards the SCVmm fabric as an Azure VmmServer resource.
 *
 * @summary Onboards the SCVmm fabric as an Azure VmmServer resource.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VmmServers_CreateOrUpdate_MinimumSet_Gen.json
 */
async function vmmServersCreateOrUpdateMinimumSet() {
  const subscriptionId =
    process.env["SCVMM_SUBSCRIPTION_ID"] || "79332E5A-630B-480F-A266-A941C015AB19";
  const resourceGroupName = process.env["SCVMM_RESOURCE_GROUP"] || "rgscvmm";
  const vmmServerName = "w";
  const resource = {
    extendedLocation: {},
    location: "hslxkyzktvwpqbypvs",
  };
  const credential = new DefaultAzureCredential();
  const client = new ScVmm(credential, subscriptionId);
  const result = await client.vmmServers.beginCreateOrUpdateAndWait(
    resourceGroupName,
    vmmServerName,
    resource,
  );
  console.log(result);
}
