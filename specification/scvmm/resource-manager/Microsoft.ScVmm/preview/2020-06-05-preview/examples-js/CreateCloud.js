const { Scvmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Onboards the ScVmm fabric cloud as an Azure cloud resource.
 *
 * @summary Onboards the ScVmm fabric cloud as an Azure cloud resource.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/CreateCloud.json
 */
async function createCloud() {
  const subscriptionId =
    process.env["SCVMM_SUBSCRIPTION_ID"] || "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = process.env["SCVMM_RESOURCE_GROUP"] || "testrg";
  const cloudName = "HRCloud";
  const body = {
    extendedLocation: {
      name: "/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.Arc/customLocations/contoso",
      type: "customLocation",
    },
    location: "East US",
    uuid: "aaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
    vmmServerId:
      "/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.SCVMM/VMMServers/ContosoVMMServer",
  };
  const credential = new DefaultAzureCredential();
  const client = new Scvmm(credential, subscriptionId);
  const result = await client.clouds.beginCreateOrUpdateAndWait(resourceGroupName, cloudName, body);
  console.log(result);
}
