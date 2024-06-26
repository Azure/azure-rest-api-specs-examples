const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update an APM.
 *
 * @summary Create or update an APM.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/Apms_CreateOrUpdate.json
 */
async function apmsCreateOrUpdate() {
  const subscriptionId =
    process.env["APPPLATFORM_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APPPLATFORM_RESOURCE_GROUP"] || "myResourceGroup";
  const serviceName = "myservice";
  const apmName = "myappinsights";
  const apmResource = {
    properties: {
      type: "ApplicationInsights",
      properties: { anyString: "any-string", samplingRate: "12.0" },
      secrets: {
        connectionString:
          "XXXXXXXXXXXXXXXXX=XXXXXXXXXXXXX-XXXXXXXXXXXXXXXXXXX;XXXXXXXXXXXXXXXXX=XXXXXXXXXXXXXXXXXXX",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.apms.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serviceName,
    apmName,
    apmResource,
  );
  console.log(result);
}
