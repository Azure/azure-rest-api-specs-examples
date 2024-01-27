const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update an VolumeGroup.
 *
 * @summary Update an VolumeGroup.
 * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/VolumeGroups_Update_MaximumSet_Gen.json
 */
async function volumeGroupsUpdateMaximumSetGen() {
  const subscriptionId = process.env["ELASTICSANS_SUBSCRIPTION_ID"] || "subscriptionid";
  const resourceGroupName = process.env["ELASTICSANS_RESOURCE_GROUP"] || "resourcegroupname";
  const elasticSanName = "elasticsanname";
  const volumeGroupName = "volumegroupname";
  const parameters = {
    identity: { type: "None", userAssignedIdentities: { key7482: {} } },
    properties: {
      encryption: "EncryptionAtRestWithPlatformKey",
      encryptionProperties: {
        encryptionIdentity: { encryptionUserAssignedIdentity: "im" },
        keyVaultProperties: {
          keyName: "sftaiernmrzypnrkpakrrawxcbsqzc",
          keyVaultUri: "https://microsoft.com/axmblwp",
          keyVersion: "c",
        },
      },
      networkAcls: {
        virtualNetworkRules: [
          {
            action: "Allow",
            virtualNetworkResourceId:
              "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}",
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
