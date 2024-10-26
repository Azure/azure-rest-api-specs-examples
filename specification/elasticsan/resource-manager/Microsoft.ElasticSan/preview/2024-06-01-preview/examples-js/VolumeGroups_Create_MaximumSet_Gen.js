const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a Volume Group.
 *
 * @summary Create a Volume Group.
 * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2024-06-01-preview/examples/VolumeGroups_Create_MaximumSet_Gen.json
 */
async function volumeGroupsCreateMaximumSetGen() {
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
  const result = await client.volumeGroups.beginCreateAndWait(
    resourceGroupName,
    elasticSanName,
    volumeGroupName,
    parameters,
  );
  console.log(result);
}
