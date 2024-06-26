const { LogicManagementClient } = require("@azure/arm-logic");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update an assembly for an integration account.
 *
 * @summary Create or update an assembly for an integration account.
 * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountAssemblies_CreateOrUpdate.json
 */
async function createOrUpdateAnAccountAssembly() {
  const subscriptionId =
    process.env["LOGIC_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["LOGIC_RESOURCE_GROUP"] || "testResourceGroup";
  const integrationAccountName = "testIntegrationAccount";
  const assemblyArtifactName = "testAssembly";
  const assemblyArtifact = {
    location: "westus",
    properties: {
      assemblyName: "System.IdentityModel.Tokens.Jwt",
      content: "Base64 encoded Assembly Content",
      metadata: {},
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new LogicManagementClient(credential, subscriptionId);
  const result = await client.integrationAccountAssemblies.createOrUpdate(
    resourceGroupName,
    integrationAccountName,
    assemblyArtifactName,
    assemblyArtifact
  );
  console.log(result);
}
