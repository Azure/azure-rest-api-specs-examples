const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Performs a connectivity check between the API Management service and a given destination, and returns metrics for the connection, as well as errors encountered while trying to establish it.
 *
 * @summary Performs a connectivity check between the API Management service and a given destination, and returns metrics for the connection, as well as errors encountered while trying to establish it.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementPerformConnectivityCheckHttpConnect.json
 */
async function httpConnectivityCheck() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const connectivityCheckRequestParams = {
    destination: { address: "https://microsoft.com", port: 3306 },
    protocolConfiguration: {
      httpConfiguration: {
        method: "GET",
        headers: [{ name: "Authorization", value: "Bearer myPreciousToken" }],
        validStatusCodes: [200, 204],
      },
    },
    source: { region: "northeurope" },
    protocol: "HTTPS",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.beginPerformConnectivityCheckAsyncAndWait(
    resourceGroupName,
    serviceName,
    connectivityCheckRequestParams,
  );
  console.log(result);
}
