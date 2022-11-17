const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete an Annotation of an Application Insights component.
 *
 * @summary Delete an Annotation of an Application Insights component.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/AnnotationsDelete.json
 */
async function annotationsDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "my-resource-group";
  const resourceName = "my-component";
  const annotationId = "bb820f1b-3110-4a8b-ba2c-8c1129d7eb6a";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.annotations.delete(resourceGroupName, resourceName, annotationId);
  console.log(result);
}

annotationsDelete().catch(console.error);
