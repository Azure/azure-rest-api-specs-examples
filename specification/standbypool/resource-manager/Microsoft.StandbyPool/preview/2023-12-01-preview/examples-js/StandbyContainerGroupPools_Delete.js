const { StandbyPoolManagementClient } = require("@azure/arm-standbypool");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a StandbyContainerGroupPoolResource
 *
 * @summary Delete a StandbyContainerGroupPoolResource
 * x-ms-original-file: specification/standbypool/resource-manager/Microsoft.StandbyPool/preview/2023-12-01-preview/examples/StandbyContainerGroupPools_Delete.json
 */
async function standbyContainerGroupPoolsDelete() {
  const subscriptionId =
    process.env["STANDBYPOOL_SUBSCRIPTION_ID"] || "8CC31D61-82D7-4B2B-B9DC-6B924DE7D229";
  const resourceGroupName = process.env["STANDBYPOOL_RESOURCE_GROUP"] || "rgstandbypool";
  const standbyContainerGroupPoolName = "pool";
  const credential = new DefaultAzureCredential();
  const client = new StandbyPoolManagementClient(credential, subscriptionId);
  const result = await client.standbyContainerGroupPools.beginDeleteAndWait(
    resourceGroupName,
    standbyContainerGroupPoolName,
  );
  console.log(result);
}
