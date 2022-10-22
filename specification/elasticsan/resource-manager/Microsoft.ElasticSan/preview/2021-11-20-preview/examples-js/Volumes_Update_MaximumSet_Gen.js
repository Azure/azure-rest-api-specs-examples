const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update an Volume.
 *
 * @summary Update an Volume.
 * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2021-11-20-preview/examples/Volumes_Update_MaximumSet_Gen.json
 */
async function volumesUpdateMaximumSetGen() {
  const subscriptionId = "aaaaaaaaaaaaaaaaaa";
  const resourceGroupName = "rgelasticsan";
  const elasticSanName = "ti7q-k952-1qB3J_5";
  const volumeGroupName = "u_5I_1j4t3";
  const volumeName = "9132y";
  const parameters = {
    sizeGiB: 9,
    tags: { key1864: "aaaaaaaaa" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ElasticSanManagement(credential, subscriptionId);
  const result = await client.volumes.beginUpdateAndWait(
    resourceGroupName,
    elasticSanName,
    volumeGroupName,
    volumeName,
    parameters
  );
  console.log(result);
}

volumesUpdateMaximumSetGen().catch(console.error);
