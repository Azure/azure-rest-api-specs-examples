const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a private workbook that has already been added.
 *
 * @summary Updates a private workbook that has already been added.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2021-03-08/examples/MyWorkbookUpdate.json
 */
async function workbookUpdate() {
  const subscriptionId = "6b643656-33eb-422f-aee8-3ac145d124af";
  const resourceGroupName = "my-resource-group";
  const resourceName = "deadb33f-5e0d-4064-8ebb-1a4ed0313eb2";
  const workbookProperties = {
    name: "deadb33f-8bee-4d3b-a059-9be8dac93960",
    category: "workbook",
    displayName: "Blah Blah Blah",
    kind: "user",
    location: "west us",
    serializedData:
      '{"version":"Notebook/1.0","items":[{"type":1,"content":"{"json":"## New workbook\\r\\n---\\r\\n\\r\\nWelcome to your new workbook.  This area will display text formatted as markdown.\\r\\n\\r\\n\\r\\nWe\'ve included a basic analytics query to get you started. Use the `Edit` button below each section to configure it or add more sections."}","halfWidth":null,"conditionalVisibility":null},{"type":3,"content":"{"version":"KqlItem/1.0","query":"union withsource=TableName *\\n| summarize Count=count() by TableName\\n| render barchart","showQuery":false,"size":1,"aggregation":0,"showAnnotations":false}","halfWidth":null,"conditionalVisibility":null}],"isLocked":false}',
    sourceId:
      "/subscriptions/00000000-0000-0000-0000-00000000/resourceGroups/MyGroup/providers/Microsoft.Web/sites/MyTestApp-CodeLens",
    storageUri: undefined,
    tags: { 0: "TagSample01", 1: "TagSample02" },
    version: "ME",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.myWorkbooks.update(
    resourceGroupName,
    resourceName,
    workbookProperties,
  );
  console.log(result);
}

workbookUpdate().catch(console.error);
