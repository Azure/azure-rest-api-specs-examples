const { AutomanageClient } = require("@azure/arm-automanage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get information about a Automanage best practice version
 *
 * @summary Get information about a Automanage best practice version
 * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/getBestPracticeVersion.json
 */
async function getAnAutomanageBestPracticeVersion() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const bestPracticeName = "azureBestPracticesProduction";
  const versionName = "version1";
  const credential = new DefaultAzureCredential();
  const client = new AutomanageClient(credential, subscriptionId);
  const result = await client.bestPracticesVersions.get(bestPracticeName, versionName);
  console.log(result);
}

getAnAutomanageBestPracticeVersion().catch(console.error);
