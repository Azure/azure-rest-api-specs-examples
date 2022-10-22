const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update an VolumeGroup.
 *
 * @summary Update an VolumeGroup.
 * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2021-11-20-preview/examples/VolumeGroups_Update_MaximumSet_Gen.json
 */
async function volumeGroupsUpdateMaximumSetGen() {
  const subscriptionId = "aaaaaaaaaaaaaaaaaa";
  const resourceGroupName = "rgelasticsan";
  const elasticSanName = "ti7q-k952-1qB3J_5";
  const volumeGroupName = "u_5I_1j4t3";
  const parameters = {
    encryption: "EncryptionAtRestWithPlatformKey",
    networkAcls: {
      virtualNetworkRules: [{ action: "Allow", virtualNetworkResourceId: "aaaaaaaaaaaaaaaa" }],
    },
    protocolType: "Iscsi",
    tags: { key7542: "aaaaaaaaaaaaaaaaaaaa" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ElasticSanManagement(credential, subscriptionId);
  const result = await client.volumeGroups.beginUpdateAndWait(
    resourceGroupName,
    elasticSanName,
    volumeGroupName,
    parameters
  );
  console.log(result);
}

volumeGroupsUpdateMaximumSetGen().catch(console.error);
