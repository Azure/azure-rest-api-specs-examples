const { ResourceManagementClient } = require("@azure/arm-resources");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Returns changes that will be made by the deployment if executed at the scope of the tenant group.
 *
 * @summary Returns changes that will be made by the deployment if executed at the scope of the tenant group.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2025-03-01/examples/PostDeploymentWhatIfOnTenant.json
 */
async function predictTemplateChangesAtManagementGroupScope() {
  const deploymentName = "exampleDeploymentName";
  const parameters = {
    location: "eastus",
    properties: {
      mode: "Incremental",
      parameters: {},
      templateLink: { uri: "https://example.com/exampleTemplate.json" },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ResourceManagementClient(credential);
  const result = await client.deployments.beginWhatIfAtTenantScopeAndWait(
    deploymentName,
    parameters,
  );
  console.log(result);
}
