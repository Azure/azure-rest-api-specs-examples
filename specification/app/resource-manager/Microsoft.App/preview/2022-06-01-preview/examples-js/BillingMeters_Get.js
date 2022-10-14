const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get all billingMeters for a location.
 *
 * @summary Get all billingMeters for a location.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-06-01-preview/examples/BillingMeters_Get.json
 */
async function billingMetersGet() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const location = "East US";
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.billingMeters.get(location);
  console.log(result);
}

billingMetersGet().catch(console.error);
