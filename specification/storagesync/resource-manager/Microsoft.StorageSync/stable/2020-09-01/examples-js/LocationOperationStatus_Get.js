const { MicrosoftStorageSync } = require("@azure/arm-storagesync");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get Operation status
 *
 * @summary Get Operation status
 * x-ms-original-file: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/LocationOperationStatus_Get.json
 */
async function workflowsGet() {
  const subscriptionId = "52b8da2f-61e0-4a1f-8dde-336911f367fb";
  const locationName = "westus";
  const operationId =
    "eyJwYXJ0aXRpb25JZCI6ImE1ZDNiMDU4LTYwN2MtNDI0Ny05Y2FmLWJlZmU4NGQ0ZDU0NyIsIndvcmtmbG93SWQiOiJjYzg1MTY2YS0xMjI2LTQ4MGYtYWM5ZC1jMmRhNTVmY2M2ODYiLCJ3b3JrZmxvd09wZXJhdGlvbklkIjoiOTdmODU5ZTAtOGY1MC00ZTg4LWJkZDEtNWZlYzgwYTVlYzM0tui=";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftStorageSync(credential, subscriptionId);
  const result = await client.locationOperationStatus(locationName, operationId);
  console.log(result);
}

workflowsGet().catch(console.error);
