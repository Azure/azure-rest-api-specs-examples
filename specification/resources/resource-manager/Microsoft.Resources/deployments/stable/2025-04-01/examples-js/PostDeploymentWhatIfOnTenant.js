const { DeploymentsClient } = require("@azure/arm-resourcesdeployments");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Returns changes that will be made by the deployment if executed at the scope of the tenant group.
 *
 * @summary Returns changes that will be made by the deployment if executed at the scope of the tenant group.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/deployments/stable/2025-04-01/examples/PostDeploymentWhatIfOnTenant.json
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
  const client = new DeploymentsClient(credential);
  const result = await client.deployments.beginWhatIfAtTenantScopeAndWait(
    deploymentName,
    parameters,
  );
  console.log(result);
}
