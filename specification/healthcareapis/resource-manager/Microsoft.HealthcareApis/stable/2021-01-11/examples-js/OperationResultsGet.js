const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

async function getOperationResult() {
  const subscriptionId = "subid";
  const locationName = "westus";
  const operationResultId = "exampleid";
  const credential = new DefaultAzureCredential();
  const client = new HealthcareApisManagementClient(credential, subscriptionId);
  const result = await client.operationResults.get(locationName, operationResultId);
  console.log(result);
}

getOperationResult().catch(console.error);
