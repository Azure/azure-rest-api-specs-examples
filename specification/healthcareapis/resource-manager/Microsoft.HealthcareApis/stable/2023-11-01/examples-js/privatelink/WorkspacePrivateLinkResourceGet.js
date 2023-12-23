const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a private link resource that need to be created for a workspace.
 *
 * @summary Gets a private link resource that need to be created for a workspace.
 * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2023-11-01/examples/privatelink/WorkspacePrivateLinkResourceGet.json
 */
async function workspacePrivateLinkResourcesGet() {
  const subscriptionId = process.env["HEALTHCAREAPIS_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["HEALTHCAREAPIS_RESOURCE_GROUP"] || "testRG";
  const workspaceName = "workspace1";
  const groupName = "healthcareworkspace";
  const credential = new DefaultAzureCredential();
  const client = new HealthcareApisManagementClient(credential, subscriptionId);
  const result = await client.workspacePrivateLinkResources.get(
    resourceGroupName,
    workspaceName,
    groupName,
  );
  console.log(result);
}
