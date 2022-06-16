const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Generates a preview workflow file for the static site
 *
 * @summary Description for Generates a preview workflow file for the static site
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GenerateStaticSiteWorkflowPreview.json
 */
async function generatesAPreviewWorkflowFileForTheStaticSite() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const location = "West US 2";
  const staticSitesWorkflowPreviewRequest = {
    branch: "master",
    buildProperties: {
      apiLocation: "api",
      appArtifactLocation: "build",
      appLocation: "app",
    },
    repositoryUrl: "https://github.com/username/RepoName",
  };
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.staticSites.previewWorkflow(
    location,
    staticSitesWorkflowPreviewRequest
  );
  console.log(result);
}

generatesAPreviewWorkflowFileForTheStaticSite().catch(console.error);
