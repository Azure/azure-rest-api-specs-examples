const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a linked service instance.
 *
 * @summary Deletes a linked service instance.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/LinkedServicesDelete.json
 */
async function linkedServicesDelete() {
  const subscriptionId =
    process.env["OPERATIONALINSIGHTS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-00000000000";
  const resourceGroupName = process.env["OPERATIONALINSIGHTS_RESOURCE_GROUP"] || "rg1";
  const workspaceName = "TestLinkWS";
  const linkedServiceName = "Cluster";
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const result = await client.linkedServices.beginDeleteAndWait(
    resourceGroupName,
    workspaceName,
    linkedServiceName
  );
  console.log(result);
}
