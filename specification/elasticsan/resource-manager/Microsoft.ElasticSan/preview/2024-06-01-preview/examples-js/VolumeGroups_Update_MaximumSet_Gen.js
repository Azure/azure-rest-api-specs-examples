const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update an VolumeGroup.
 *
 * @summary Update an VolumeGroup.
 * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2024-06-01-preview/examples/VolumeGroups_Update_MaximumSet_Gen.json
 */
async function volumeGroupsUpdateMaximumSetGen() {
  const subscriptionId = process.env["ELASTICSANS_SUBSCRIPTION_ID"] || "subscriptionid";
  const resourceGroupName = process.env["ELASTICSANS_RESOURCE_GROUP"] || "resourcegroupname";
  const elasticSanName = "elasticsanname";
  const volumeGroupName = "volumegroupname";
  const parameters = {
    identity: { type: "None", userAssignedIdentities: { key1006: {} } },
    properties: {
      encryption: "EncryptionAtRestWithPlatformKey",
      encryptionProperties: {
        encryptionIdentity: {
          encryptionUserAssignedIdentity: "gfhkfbozahmmwluqndfgxunssafa",
        },
        keyVaultProperties: {
          keyName: "lunpapamzeimppgobraxjt",
          keyVaultUri: "https://microsoft.com/a",
          keyVersion: "oemygbnfmqhijmonkqfqmy",
        },
      },
      enforceDataIntegrityCheckForIscsi: true,
      networkAcls: {
        virtualNetworkRules: [
          {
            action: "Allow",
            virtualNetworkResourceId: "bkhwaiqvvaguymsmnzzbzz",
          },
        ],
      },
      protocolType: "Iscsi",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ElasticSanManagement(credential, subscriptionId);
  const result = await client.volumeGroups.beginUpdateAndWait(
    resourceGroupName,
    elasticSanName,
    volumeGroupName,
    parameters,
  );
  console.log(result);
}
