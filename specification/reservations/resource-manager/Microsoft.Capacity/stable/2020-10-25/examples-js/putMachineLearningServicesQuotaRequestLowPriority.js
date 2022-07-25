const { AzureReservationAPI } = require("@azure/arm-reservations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update the quota (service limits) of a resource to the requested value.
 Steps:

  1. Make the Get request to get the quota information for specific resource.

  2. To increase the quota, update the limit field in the response from Get request to new value.

  3. Submit the JSON to the quota request API to update the quota.
  The Create quota request may be constructed as follows. The PUT operation can be used to update the quota.
 *
 * @summary Create or update the quota (service limits) of a resource to the requested value.
 Steps:

  1. Make the Get request to get the quota information for specific resource.

  2. To increase the quota, update the limit field in the response from Get request to new value.

  3. Submit the JSON to the quota request API to update the quota.
  The Create quota request may be constructed as follows. The PUT operation can be used to update the quota.
 * x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2020-10-25/examples/putMachineLearningServicesQuotaRequestLowPriority.json
 */
async function quotasRequestPutForMachineLearningServicesLowPriorityResource() {
  const subscriptionId = "D7EC67B3-7657-4966-BFFC-41EFD36BAAB3";
  const providerId = "Microsoft.MachineLearningServices";
  const location = "eastus";
  const resourceName = "TotalLowPriorityCores";
  const createQuotaRequest = {
    properties: {
      name: { value: "TotalLowPriorityCores" },
      limit: 200,
      resourceType: "lowPriority",
      unit: "Count",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureReservationAPI(credential);
  const result = await client.quota.beginCreateOrUpdateAndWait(
    subscriptionId,
    providerId,
    location,
    resourceName,
    createQuotaRequest
  );
  console.log(result);
}

quotasRequestPutForMachineLearningServicesLowPriorityResource().catch(console.error);
