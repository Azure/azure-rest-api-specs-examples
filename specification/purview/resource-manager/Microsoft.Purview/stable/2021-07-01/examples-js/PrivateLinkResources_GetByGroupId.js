const { PurviewManagementClient } = require("@azure/arm-purview");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a privately linkable resources for an account with given group identifier
 *
 * @summary Gets a privately linkable resources for an account with given group identifier
 * x-ms-original-file: specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/PrivateLinkResources_GetByGroupId.json
 */
async function privateLinkResourcesGetByGroupId() {
  const subscriptionId = "12345678-1234-1234-12345678abc";
  const resourceGroupName = "SampleResourceGroup";
  const accountName = "account1";
  const groupId = "group1";
  const credential = new DefaultAzureCredential();
  const client = new PurviewManagementClient(credential, subscriptionId);
  const result = await client.privateLinkResources.getByGroupId(
    resourceGroupName,
    accountName,
    groupId
  );
  console.log(result);
}

privateLinkResourcesGetByGroupId().catch(console.error);
