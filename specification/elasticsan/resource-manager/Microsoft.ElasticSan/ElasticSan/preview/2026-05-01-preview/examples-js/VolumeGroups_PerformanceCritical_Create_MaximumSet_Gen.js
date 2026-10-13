const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a Volume Group.
 *
 * @summary create a Volume Group.
 * x-ms-original-file: 2026-05-01-preview/VolumeGroups_PerformanceCritical_Create_MaximumSet_Gen.json
 */
async function volumeGroupsPerformanceCriticalCreateMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscriptionid";
  const client = new ElasticSanManagement(credential, subscriptionId);
  const result = await client.volumeGroups.create(
    "resourcegroupname",
    "elasticsanname",
    "volumegroupname",
    {
      identity: { type: "None", userAssignedIdentities: { key2350: {} } },
      properties: {
        encryption: "EncryptionAtRestWithPlatformKey",
        encryptionProperties: {
          encryptionIdentity: { encryptionUserAssignedIdentity: "vgbeephfgecgg" },
          keyVaultProperties: {
            keyName: "rommjwp",
            keyVaultUri: "https://microsoft.com/at",
            keyVersion: "ulmxxgzgsuhalwesmhfslq",
          },
        },
        protocolType: "DirectAttach",
        qualityOfService: "PerformanceCritical",
        reservedIops: 10000,
        reservedMBps: 800,
      },
    },
  );
  console.log(result);
}
