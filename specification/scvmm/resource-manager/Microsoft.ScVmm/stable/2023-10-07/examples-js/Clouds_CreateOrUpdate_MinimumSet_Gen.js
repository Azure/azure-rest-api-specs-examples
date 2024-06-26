const { ScVmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Onboards the ScVmm fabric cloud as an Azure cloud resource.
 *
 * @summary Onboards the ScVmm fabric cloud as an Azure cloud resource.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/Clouds_CreateOrUpdate_MinimumSet_Gen.json
 */
async function cloudsCreateOrUpdateMinimumSet() {
  const subscriptionId =
    process.env["SCVMM_SUBSCRIPTION_ID"] || "79332E5A-630B-480F-A266-A941C015AB19";
  const resourceGroupName = process.env["SCVMM_RESOURCE_GROUP"] || "rgscvmm";
  const cloudResourceName = "-";
  const resource = { extendedLocation: {}, location: "khwsdmaxfhmbu" };
  const credential = new DefaultAzureCredential();
  const client = new ScVmm(credential, subscriptionId);
  const result = await client.clouds.beginCreateOrUpdateAndWait(
    resourceGroupName,
    cloudResourceName,
    resource,
  );
  console.log(result);
}
