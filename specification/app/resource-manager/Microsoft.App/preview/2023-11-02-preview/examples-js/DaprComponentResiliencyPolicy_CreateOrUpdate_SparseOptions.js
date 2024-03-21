const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a resiliency policy for a Dapr component.
 *
 * @summary Creates or updates a resiliency policy for a Dapr component.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/DaprComponentResiliencyPolicy_CreateOrUpdate_SparseOptions.json
 */
async function createOrUpdateDaprComponentResiliencyPolicyWithSparseOptions() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "8efdecc5-919e-44eb-b179-915dca89ebf9";
  const resourceGroupName = process.env["APPCONTAINERS_RESOURCE_GROUP"] || "examplerg";
  const environmentName = "myenvironment";
  const componentName = "mydaprcomponent";
  const name = "myresiliencypolicy";
  const daprComponentResiliencyPolicyEnvelope = {
    inboundPolicy: {
      circuitBreakerPolicy: { consecutiveErrors: 3, timeoutInSeconds: 20 },
      httpRetryPolicy: {
        maxRetries: 5,
        retryBackOff: {
          initialDelayInMilliseconds: 2000,
          maxIntervalInMilliseconds: 5500,
        },
      },
    },
    outboundPolicy: { timeoutPolicy: { responseTimeoutInSeconds: 12 } },
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.daprComponentResiliencyPolicies.createOrUpdate(
    resourceGroupName,
    environmentName,
    componentName,
    name,
    daprComponentResiliencyPolicyEnvelope,
  );
  console.log(result);
}
