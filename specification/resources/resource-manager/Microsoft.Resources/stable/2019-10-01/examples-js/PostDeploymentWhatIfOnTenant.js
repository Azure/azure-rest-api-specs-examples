const { ResourceManagementClient } = require("@azure/arm-resources-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns changes that will be made by the deployment if executed at the scope of the tenant group.
 *
 * @summary Returns changes that will be made by the deployment if executed at the scope of the tenant group.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2019-10-01/examples/PostDeploymentWhatIfOnTenant.json
 */
async function predictTemplateChangesAtManagementGroupScope() {
  const subscriptionId =
    process.env["RESOURCES_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
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
  const client = new ResourceManagementClient(credential, subscriptionId);
  const result = await client.deployments.beginWhatIfAtTenantScopeAndWait(
    deploymentName,
    parameters
  );
  console.log(result);
}
