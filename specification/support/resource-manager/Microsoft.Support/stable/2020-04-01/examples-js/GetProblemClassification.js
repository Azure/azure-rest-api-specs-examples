const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

async function getsDetailsOfProblemClassificationForAzureService() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const serviceName = "service_guid";
  const problemClassificationName = "problemClassification_guid";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential, subscriptionId);
  const result = await client.problemClassifications.get(serviceName, problemClassificationName);
  console.log(result);
}

getsDetailsOfProblemClassificationForAzureService().catch(console.error);
