const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a Volume Group.
 *
 * @summary Create a Volume Group.
 * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/VolumeGroups_Create_MaximumSet_Gen.json
 */
async function volumeGroupsCreateMaximumSetGen() {
  const subscriptionId = process.env["ELASTICSANS_SUBSCRIPTION_ID"] || "subscriptionid";
  const resourceGroupName = process.env["ELASTICSANS_RESOURCE_GROUP"] || "resourcegroupname";
  const elasticSanName = "elasticsanname";
  const volumeGroupName = "volumegroupname";
  const parameters = {
    identity: { type: "None", userAssignedIdentities: { key7482: {} } },
    properties: {
      encryption: "EncryptionAtRestWithCustomerManagedKey",
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
  const result = await client.volumeGroups.beginCreateAndWait(
    resourceGroupName,
    elasticSanName,
    volumeGroupName,
    parameters,
  );
  console.log(result);
}
