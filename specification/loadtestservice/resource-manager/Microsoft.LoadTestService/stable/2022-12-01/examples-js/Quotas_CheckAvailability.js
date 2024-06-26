const { LoadTestClient } = require("@azure/arm-loadtesting");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Check Quota Availability on quota bucket per region per subscription.
 *
 * @summary Check Quota Availability on quota bucket per region per subscription.
 * x-ms-original-file: specification/loadtestservice/resource-manager/Microsoft.LoadTestService/stable/2022-12-01/examples/Quotas_CheckAvailability.json
 */
async function quotasCheckAvailability() {
  const subscriptionId =
    process.env["LOADTESTSERVICE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const location = "westus";
  const quotaBucketName = "testQuotaBucket";
  const quotaBucketRequest = {
    currentQuota: 40,
    currentUsage: 20,
    dimensions: { location: "westus", subscriptionId: "testsubscriptionId" },
    newQuota: 50,
  };
  const credential = new DefaultAzureCredential();
  const client = new LoadTestClient(credential, subscriptionId);
  const result = await client.quotas.checkAvailability(
    location,
    quotaBucketName,
    quotaBucketRequest
  );
  console.log(result);
}
