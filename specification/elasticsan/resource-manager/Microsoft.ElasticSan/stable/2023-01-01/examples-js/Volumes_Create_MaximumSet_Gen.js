const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a Volume.
 *
 * @summary Create a Volume.
 * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/Volumes_Create_MaximumSet_Gen.json
 */
async function volumesCreateMaximumSetGen() {
  const subscriptionId = process.env["ELASTICSANS_SUBSCRIPTION_ID"] || "subscriptionid";
  const resourceGroupName = process.env["ELASTICSANS_RESOURCE_GROUP"] || "resourcegroupname";
  const elasticSanName = "elasticsanname";
  const volumeGroupName = "volumegroupname";
  const volumeName = "volumename";
  const parameters = {
    properties: {
      creationData: { createSource: "None", sourceId: "ARM Id of Resource" },
      managedBy: { resourceId: "mtkeip" },
      sizeGiB: 9,
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ElasticSanManagement(credential, subscriptionId);
  const result = await client.volumes.beginCreateAndWait(
    resourceGroupName,
    elasticSanName,
    volumeGroupName,
    volumeName,
    parameters,
  );
  console.log(result);
}
