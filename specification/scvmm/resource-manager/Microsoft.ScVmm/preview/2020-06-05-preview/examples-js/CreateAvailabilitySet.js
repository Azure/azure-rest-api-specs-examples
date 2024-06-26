const { Scvmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Onboards the ScVmm availability set as an Azure resource.
 *
 * @summary Onboards the ScVmm availability set as an Azure resource.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/CreateAvailabilitySet.json
 */
async function createAvailabilitySet() {
  const subscriptionId =
    process.env["SCVMM_SUBSCRIPTION_ID"] || "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = process.env["SCVMM_RESOURCE_GROUP"] || "testrg";
  const availabilitySetName = "HRAvailabilitySet";
  const body = {
    availabilitySetName: "hr-avset",
    extendedLocation: {
      name: "/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.Arc/customLocations/contoso",
      type: "customLocation",
    },
    location: "East US",
    vmmServerId:
      "/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ScVmm/VMMServers/ContosoVMMServer",
  };
  const credential = new DefaultAzureCredential();
  const client = new Scvmm(credential, subscriptionId);
  const result = await client.availabilitySets.beginCreateOrUpdateAndWait(
    resourceGroupName,
    availabilitySetName,
    body
  );
  console.log(result);
}
