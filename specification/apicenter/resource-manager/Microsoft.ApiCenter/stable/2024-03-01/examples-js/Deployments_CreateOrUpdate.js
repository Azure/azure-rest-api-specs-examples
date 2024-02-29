const { AzureAPICenter } = require("@azure/arm-apicenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates new or updates existing API deployment.
 *
 * @summary Creates new or updates existing API deployment.
 * x-ms-original-file: specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/Deployments_CreateOrUpdate.json
 */
async function deploymentsCreateOrUpdate() {
  const subscriptionId =
    process.env["APICENTER_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APICENTER_RESOURCE_GROUP"] || "contoso-resources";
  const serviceName = "contoso";
  const workspaceName = "default";
  const apiName = "echo-api";
  const deploymentName = "production";
  const payload = {
    properties: {
      description: "Public cloud production deployment.",
      definitionId: "/workspaces/default/apis/echo-api/versions/2023-01-01/definitions/openapi",
      environmentId: "/workspaces/default/environments/production",
      server: { runtimeUri: ["https://api.contoso.com"] },
      state: "active",
      title: "Production deployment",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureAPICenter(credential, subscriptionId);
  const result = await client.deployments.createOrUpdate(
    resourceGroupName,
    serviceName,
    workspaceName,
    apiName,
    deploymentName,
    payload,
  );
  console.log(result);
}
