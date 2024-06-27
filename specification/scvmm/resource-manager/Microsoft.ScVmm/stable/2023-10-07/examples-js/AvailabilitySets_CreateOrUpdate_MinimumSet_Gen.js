const { ScVmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Onboards the ScVmm availability set as an Azure resource.
 *
 * @summary Onboards the ScVmm availability set as an Azure resource.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/AvailabilitySets_CreateOrUpdate_MinimumSet_Gen.json
 */
async function availabilitySetsCreateOrUpdateMinimumSet() {
  const subscriptionId =
    process.env["SCVMM_SUBSCRIPTION_ID"] || "79332E5A-630B-480F-A266-A941C015AB19";
  const resourceGroupName = process.env["SCVMM_RESOURCE_GROUP"] || "rgscvmm";
  const availabilitySetResourceName = "_";
  const resource = {
    extendedLocation: {},
    location: "jelevilan",
  };
  const credential = new DefaultAzureCredential();
  const client = new ScVmm(credential, subscriptionId);
  const result = await client.availabilitySets.beginCreateOrUpdateAndWait(
    resourceGroupName,
    availabilitySetResourceName,
    resource,
  );
  console.log(result);
}
