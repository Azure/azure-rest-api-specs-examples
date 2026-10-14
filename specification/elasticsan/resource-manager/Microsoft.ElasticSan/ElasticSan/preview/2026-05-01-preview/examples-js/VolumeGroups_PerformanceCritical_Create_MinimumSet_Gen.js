const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a Volume Group.
 *
 * @summary create a Volume Group.
 * x-ms-original-file: 2026-05-01-preview/VolumeGroups_PerformanceCritical_Create_MinimumSet_Gen.json
 */
async function volumeGroupsPerformanceCriticalCreateMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscriptionid";
  const client = new ElasticSanManagement(credential, subscriptionId);
  const result = await client.volumeGroups.create(
    "resourcegroupname",
    "elasticsanname",
    "volumegroupname",
    {
      properties: {
        protocolType: "DirectAttach",
        qualityOfService: "PerformanceCritical",
        reservedIops: 10000,
        reservedMBps: 800,
      },
    },
  );
  console.log(result);
}
