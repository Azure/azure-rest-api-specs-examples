const { AzureReservationAPI } = require("@azure/arm-reservations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update the quota (service limits) of a resource to the requested value.
 * Steps:
 * 1. Make the Get request to get the quota information for specific resource.
 * 2. To increase the quota, update the limit field in the response from Get request to new value.
 * 3. Submit the JSON to the quota request API to update the quota.
 * The Create quota request may be constructed as follows. The PUT operation can be used to update the quota.
 *
 * @summary create or update the quota (service limits) of a resource to the requested value.
 * Steps:
 * 1. Make the Get request to get the quota information for specific resource.
 * 2. To increase the quota, update the limit field in the response from Get request to new value.
 * 3. Submit the JSON to the quota request API to update the quota.
 * The Create quota request may be constructed as follows. The PUT operation can be used to update the quota.
 * x-ms-original-file: 2020-10-25/putComputeOneSkuQuotaRequest.json
 */
async function quotasRequestPutForCompute() {
  const credential = new DefaultAzureCredential();
  const client = new AzureReservationAPI(credential);
  const result = await client.quota.createOrUpdate(
    "D7EC67B3-7657-4966-BFFC-41EFD36BAAB3",
    "Microsoft.Compute",
    "eastus",
    "standardFSv2Family",
    { properties: { name: { value: "standardFSv2Family" }, limit: 200, unit: "Count" } },
  );
  console.log(result);
}
