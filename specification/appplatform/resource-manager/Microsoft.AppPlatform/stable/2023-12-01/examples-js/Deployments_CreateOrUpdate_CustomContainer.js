const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a new Deployment or update an exiting Deployment.
 *
 * @summary Create a new Deployment or update an exiting Deployment.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/Deployments_CreateOrUpdate_CustomContainer.json
 */
async function deploymentsCreateOrUpdateCustomContainer() {
  const subscriptionId =
    process.env["APPPLATFORM_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APPPLATFORM_RESOURCE_GROUP"] || "myResourceGroup";
  const serviceName = "myservice";
  const appName = "myapp";
  const deploymentName = "mydeployment";
  const deploymentResource = {
    properties: {
      deploymentSettings: {
        environmentVariables: { env: "test" },
        livenessProbe: {
          disableProbe: false,
          failureThreshold: 3,
          initialDelaySeconds: 30,
          periodSeconds: 10,
          probeAction: {
            type: "HTTPGetAction",
            path: "/health",
            scheme: "HTTP",
          },
        },
        readinessProbe: {
          disableProbe: false,
          failureThreshold: 3,
          initialDelaySeconds: 30,
          periodSeconds: 10,
          probeAction: {
            type: "HTTPGetAction",
            path: "/health",
            scheme: "HTTP",
          },
        },
        resourceRequests: { cpu: "1000m", memory: "3Gi" },
        startupProbe: {
          disableProbe: false,
        },
        terminationGracePeriodSeconds: 30,
      },
      instances: [],
      source: {
        type: "Container",
        customContainer: {
          args: ["-c", "while true; do echo hello; sleep 10;done"],
          command: ["/bin/sh"],
          containerImage: "myContainerImage:v1",
          imageRegistryCredential: {
            password: "myPassword",
            username: "myUsername",
          },
          languageFramework: "springboot",
          server: "myacr.azurecr.io",
        },
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.deployments.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serviceName,
    appName,
    deploymentName,
    deploymentResource,
  );
  console.log(result);
}
