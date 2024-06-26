const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update custom domain of one lifecycle application.
 *
 * @summary Update custom domain of one lifecycle application.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/CustomDomains_Update.json
 */
async function customDomainsUpdate() {
  const subscriptionId =
    process.env["APPPLATFORM_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APPPLATFORM_RESOURCE_GROUP"] || "myResourceGroup";
  const serviceName = "myservice";
  const appName = "myapp";
  const domainName = "mydomain.com";
  const domainResource = {
    properties: {
      certName: "mycert",
      thumbprint: "934367bf1c97033f877db0f15cb1b586957d3133",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.customDomains.beginUpdateAndWait(
    resourceGroupName,
    serviceName,
    appName,
    domainName,
    domainResource,
  );
  console.log(result);
}
