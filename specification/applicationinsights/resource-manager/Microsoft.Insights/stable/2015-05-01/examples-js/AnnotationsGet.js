const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the annotation for given id.
 *
 * @summary Get the annotation for given id.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/AnnotationsGet.json
 */
async function annotationsGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "my-resource-group";
  const resourceName = "my-component";
  const annotationId = "444e2c08-274a-4bbb-a89e-d77bb720f44a";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.annotations.get(resourceGroupName, resourceName, annotationId);
  console.log(result);
}

annotationsGet().catch(console.error);
