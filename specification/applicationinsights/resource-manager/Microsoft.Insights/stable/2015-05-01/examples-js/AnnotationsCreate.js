const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Create an Annotation of an Application Insights component.
 *
 * @summary Create an Annotation of an Application Insights component.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/AnnotationsCreate.json
 */
async function annotationsCreate() {
  const subscriptionId = process.env["APPLICATIONINSIGHTS_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName =
    process.env["APPLICATIONINSIGHTS_RESOURCE_GROUP"] || "my-resource-group";
  const resourceName = "my-component";
  const annotationProperties = {
    annotationName: "TestAnnotation",
    category: "Text",
    eventTime: new Date("2018-01-31T13:41:38.657Z"),
    id: "444e2c08-274a-4bbb-a89e-d77bb720f44a",
    properties: '{"Comments":"Testing","Label":"Success"}',
  };
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.annotations.create(
    resourceGroupName,
    resourceName,
    annotationProperties,
  );
  console.log(result);
}
