const { DeveloperHubServiceClient } = require("@azure/arm-devhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Generate preview dockerfile and manifests.
 *
 * @summary Generate preview dockerfile and manifests.
 * x-ms-original-file: specification/developerhub/resource-manager/Microsoft.DevHub/preview/2022-10-11-preview/examples/GeneratePreviewArtifacts.json
 */
async function artifactGenerationProperties() {
  const subscriptionId = process.env["DEVHUB_SUBSCRIPTION_ID"] || "subscriptionId1";
  const location = "location1";
  const parameters = {
    appName: "my-app",
    dockerfileGenerationMode: "enabled",
    dockerfileOutputDirectory: "./",
    generationLanguage: "javascript",
    imageName: "myimage",
    imageTag: "latest",
    languageVersion: "14",
    manifestGenerationMode: "enabled",
    manifestOutputDirectory: "./",
    manifestType: "kube",
    namespace: "my-namespace",
    port: "80",
  };
  const credential = new DefaultAzureCredential();
  const client = new DeveloperHubServiceClient(credential, subscriptionId);
  const result = await client.generatePreviewArtifacts(location, parameters);
  console.log(result);
}
