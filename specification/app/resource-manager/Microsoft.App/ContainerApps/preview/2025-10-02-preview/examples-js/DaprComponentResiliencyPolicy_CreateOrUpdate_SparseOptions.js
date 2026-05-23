const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a resiliency policy for a Dapr component.
 *
 * @summary creates or updates a resiliency policy for a Dapr component.
 * x-ms-original-file: 2025-10-02-preview/DaprComponentResiliencyPolicy_CreateOrUpdate_SparseOptions.json
 */
async function createOrUpdateDaprComponentResiliencyPolicyWithSparseOptions() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "8efdecc5-919e-44eb-b179-915dca89ebf9";
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.daprComponentResiliencyPolicies.createOrUpdate(
    "examplerg",
    "myenvironment",
    "mydaprcomponent",
    "myresiliencypolicy",
    {
      inboundPolicy: {
        circuitBreakerPolicy: { consecutiveErrors: 3, timeoutInSeconds: 20 },
        httpRetryPolicy: {
          maxRetries: 5,
          retryBackOff: { initialDelayInMilliseconds: 2000, maxIntervalInMilliseconds: 5500 },
        },
      },
      outboundPolicy: { timeoutPolicy: { responseTimeoutInSeconds: 12 } },
    },
  );
  console.log(result);
}
