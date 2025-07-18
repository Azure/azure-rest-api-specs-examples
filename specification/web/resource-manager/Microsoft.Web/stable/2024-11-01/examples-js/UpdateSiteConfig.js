const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Description for Updates the configuration of an app.
 *
 * @summary Description for Updates the configuration of an app.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/UpdateSiteConfig.json
 */
async function updateSiteConfig() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPSERVICE_RESOURCE_GROUP"] || "testrg123";
  const name = "sitef6141";
  const siteConfig = {
    acrUseManagedIdentityCreds: false,
    alwaysOn: false,
    appCommandLine: "",
    autoHealEnabled: false,
    azureStorageAccounts: {},
    defaultDocuments: [
      "Default.htm",
      "Default.html",
      "Default.asp",
      "index.htm",
      "index.html",
      "iisstart.htm",
      "default.aspx",
      "index.php",
      "hostingstart.html",
    ],
    detailedErrorLoggingEnabled: false,
    ftpsState: "AllAllowed",
    functionAppScaleLimit: 0,
    functionsRuntimeScaleMonitoringEnabled: false,
    http20Enabled: false,
    httpLoggingEnabled: false,
    linuxFxVersion: "",
    loadBalancing: "LeastRequests",
    logsDirectorySizeLimit: 35,
    managedPipelineMode: "Integrated",
    minTlsVersion: "1.2",
    minimumElasticInstanceCount: 0,
    netFrameworkVersion: "v4.0",
    nodeVersion: "",
    numberOfWorkers: 1,
    phpVersion: "5.6",
    powerShellVersion: "",
    pythonVersion: "",
    remoteDebuggingEnabled: false,
    requestTracingEnabled: false,
    scmMinTlsVersion: "1.2",
    use32BitWorkerProcess: true,
    virtualApplications: [
      {
        physicalPath: "site\\wwwroot",
        preloadEnabled: false,
        virtualPath: "/",
      },
    ],
    vnetName: "",
    vnetPrivatePortsCount: 0,
    vnetRouteAllEnabled: false,
    webSocketsEnabled: false,
  };
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.webApps.createOrUpdateConfiguration(
    resourceGroupName,
    name,
    siteConfig,
  );
  console.log(result);
}
