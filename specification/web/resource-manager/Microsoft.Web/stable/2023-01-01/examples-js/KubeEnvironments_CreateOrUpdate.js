const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Creates or updates a Kubernetes Environment.
 *
 * @summary Description for Creates or updates a Kubernetes Environment.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/KubeEnvironments_CreateOrUpdate.json
 */
async function createKubeEnvironments() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPSERVICE_RESOURCE_GROUP"] || "examplerg";
  const name = "testkubeenv";
  const kubeEnvironmentEnvelope = {
    location: "East US",
    staticIp: "1.2.3.4",
  };
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.kubeEnvironments.beginCreateOrUpdateAndWait(
    resourceGroupName,
    name,
    kubeEnvironmentEnvelope
  );
  console.log(result);
}
