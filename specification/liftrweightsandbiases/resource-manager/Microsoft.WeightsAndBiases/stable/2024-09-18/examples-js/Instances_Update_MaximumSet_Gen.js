const { WeightsAndBiasesClient } = require("@azure/arm-weightsandbiases");
const { DefaultAzureCredential } = require("@azure/identity");

async function instancesUpdateGeneratedByMaximumSetRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "0BCB047F-CB71-4DFE-B0BD-88519F411C2F";
  const client = new WeightsAndBiasesClient(credential, subscriptionId);
  const result = await client.instances.update("rgopenapi", "myinstance", {
    tags: {},
    identity: { type: "None", userAssignedIdentities: {} },
  });
  console.log(result);
}
