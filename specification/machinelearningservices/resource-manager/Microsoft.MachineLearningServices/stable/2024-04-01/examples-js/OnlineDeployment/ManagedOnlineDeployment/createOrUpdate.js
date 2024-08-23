const { AzureMachineLearningServicesManagementClient } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update Inference Endpoint Deployment (asynchronous).
 *
 * @summary Create or update Inference Endpoint Deployment (asynchronous).
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/OnlineDeployment/ManagedOnlineDeployment/createOrUpdate.json
 */
async function createOrUpdateManagedOnlineDeployment() {
  const subscriptionId =
    process.env["MACHINELEARNING_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["MACHINELEARNING_RESOURCE_GROUP"] || "test-rg";
  const workspaceName = "my-aml-workspace";
  const endpointName = "testEndpointName";
  const deploymentName = "testDeploymentName";
  const body = {
    identity: {
      type: "SystemAssigned",
      userAssignedIdentities: { string: {} },
    },
    kind: "string",
    location: "string",
    properties: {
      description: "string",
      appInsightsEnabled: false,
      codeConfiguration: { codeId: "string", scoringScript: "string" },
      endpointComputeType: "Managed",
      environmentId: "string",
      environmentVariables: { string: "string" },
      instanceType: "string",
      livenessProbe: {
        failureThreshold: 1,
        initialDelay: "PT5M",
        period: "PT5M",
        successThreshold: 1,
        timeout: "PT5M",
      },
      model: "string",
      modelMountPath: "string",
      properties: { string: "string" },
      readinessProbe: {
        failureThreshold: 30,
        initialDelay: "PT1S",
        period: "PT10S",
        successThreshold: 1,
        timeout: "PT2S",
      },
      requestSettings: {
        maxConcurrentRequestsPerInstance: 1,
        maxQueueWait: "PT5M",
        requestTimeout: "PT5M",
      },
      scaleSettings: { scaleType: "Default" },
    },
    sku: {
      name: "string",
      capacity: 1,
      family: "string",
      size: "string",
      tier: "Free",
    },
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningServicesManagementClient(credential, subscriptionId);
  const result = await client.onlineDeployments.beginCreateOrUpdateAndWait(
    resourceGroupName,
    workspaceName,
    endpointName,
    deploymentName,
    body,
  );
  console.log(result);
}
