const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or replace an existing secret. This operation can take a while to complete.
 *
 * @summary Create or replace an existing secret. This operation can take a while to complete.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Secrets_CreateOrUpdate.json
 */
async function secretsCreateOrUpdate() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const labName = "{labName}";
  const userName = "{userName}";
  const name = "{secretName}";
  const secret = { value: "{secret}" };
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const result = await client.secrets.beginCreateOrUpdateAndWait(
    resourceGroupName,
    labName,
    userName,
    name,
    secret
  );
  console.log(result);
}

secretsCreateOrUpdate().catch(console.error);
