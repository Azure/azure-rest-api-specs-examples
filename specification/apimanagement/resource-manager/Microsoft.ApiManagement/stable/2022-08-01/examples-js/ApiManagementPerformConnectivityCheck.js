const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Performs a connectivity check between the API Management service and a given destination, and returns metrics for the connection, as well as errors encountered while trying to establish it.
 *
 * @summary Performs a connectivity check between the API Management service and a given destination, and returns metrics for the connection, as well as errors encountered while trying to establish it.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementPerformConnectivityCheck.json
 */
async function tcpConnectivityCheck() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const connectivityCheckRequestParams = {
    destination: { address: "8.8.8.8", port: 53 },
    preferredIPVersion: "IPv4",
    source: { region: "northeurope" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.beginPerformConnectivityCheckAsyncAndWait(
    resourceGroupName,
    serviceName,
    connectivityCheckRequestParams
  );
  console.log(result);
}
