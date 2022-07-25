const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a source control.
 *
 * @summary Creates a source control.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-07-01-preview/examples/sourcecontrols/CreateSourceControl.json
 */
async function createsASourceControl() {
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const resourceGroupName = "myRg";
  const workspaceName = "myWorkspace";
  const sourceControlId = "789e0c1f-4a3d-43ad-809c-e713b677b04a";
  const sourceControl = {
    description: "This is a source control",
    contentTypes: ["AnalyticRules", "Workbook"],
    displayName: "My Source Control",
    etag: '"0300bf09-0000-0000-0000-5c37296e0000"',
    repoType: "Github",
    repository: {
      branch: "master",
      displayUrl: "https://github.com/user/repo",
      pathMapping: [
        { path: "path/to/rules", contentType: "AnalyticRules" },
        { path: "path/to/workbooks", contentType: "Workbook" },
      ],
      url: "https://github.com/user/repo",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.sourceControls.create(
    resourceGroupName,
    workspaceName,
    sourceControlId,
    sourceControl
  );
  console.log(result);
}

createsASourceControl().catch(console.error);
