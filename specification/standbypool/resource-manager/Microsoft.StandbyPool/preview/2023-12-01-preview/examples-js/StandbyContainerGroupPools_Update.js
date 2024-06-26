const { StandbyPoolManagementClient } = require("@azure/arm-standbypool");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a StandbyContainerGroupPoolResource
 *
 * @summary Update a StandbyContainerGroupPoolResource
 * x-ms-original-file: specification/standbypool/resource-manager/Microsoft.StandbyPool/preview/2023-12-01-preview/examples/StandbyContainerGroupPools_Update.json
 */
async function standbyContainerGroupPoolsUpdate() {
  const subscriptionId =
    process.env["STANDBYPOOL_SUBSCRIPTION_ID"] || "8CC31D61-82D7-4B2B-B9DC-6B924DE7D229";
  const resourceGroupName = process.env["STANDBYPOOL_RESOURCE_GROUP"] || "rgstandbypool";
  const standbyContainerGroupPoolName = "pool";
  const properties = {
    properties: {
      containerGroupProperties: {
        containerGroupProfile: {
          id: "/subscriptions/8CC31D61-82D7-4B2B-B9DC-6B924DE7D229/resourceGroups/rgstandbypool/providers/Microsoft.ContainerInstance/containerGroupProfiles/cgProfile",
          revision: 2,
        },
        subnetIds: [
          {
            id: "/subscriptions/8cf6c1b6-c80f-437c-87ad-45fbaff54f73/resourceGroups/rgstandbypool/providers/Microsoft.Network/virtualNetworks/cgSubnet/subnets/cgSubnet",
          },
        ],
      },
      elasticityProfile: { maxReadyCapacity: 1743, refillPolicy: "always" },
    },
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new StandbyPoolManagementClient(credential, subscriptionId);
  const result = await client.standbyContainerGroupPools.update(
    resourceGroupName,
    standbyContainerGroupPoolName,
    properties,
  );
  console.log(result);
}
