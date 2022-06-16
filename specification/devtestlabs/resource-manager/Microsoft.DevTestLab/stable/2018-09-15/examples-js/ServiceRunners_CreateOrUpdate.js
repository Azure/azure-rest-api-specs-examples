const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or replace an existing service runner.
 *
 * @summary Create or replace an existing service runner.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ServiceRunners_CreateOrUpdate.json
 */
async function serviceRunnersCreateOrUpdate() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const labName = "{devtestlabName}";
  const name = "{servicerunnerName}";
  const serviceRunner = {
    identity: {
      type: "{identityType}",
      clientSecretUrl: "{identityClientSecretUrl}",
      principalId: "{identityPrincipalId}",
      tenantId: "{identityTenantId}",
    },
    location: "{location}",
    tags: { tagName1: "tagValue1" },
  };
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const result = await client.serviceRunners.createOrUpdate(
    resourceGroupName,
    labName,
    name,
    serviceRunner
  );
  console.log(result);
}

serviceRunnersCreateOrUpdate().catch(console.error);
