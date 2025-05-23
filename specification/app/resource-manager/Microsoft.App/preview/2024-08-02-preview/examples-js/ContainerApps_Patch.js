const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Patches a Container App using JSON Merge Patch
 *
 * @summary Patches a Container App using JSON Merge Patch
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/ContainerApps_Patch.json
 */
async function patchContainerApp() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPCONTAINERS_RESOURCE_GROUP"] || "rg";
  const containerAppName = "testcontainerApp0";
  const containerAppEnvelope = {
    configuration: {
      dapr: {
        appPort: 3000,
        appProtocol: "http",
        enableApiLogging: true,
        enabled: true,
        httpMaxRequestSize: 10,
        httpReadBufferSize: 30,
        logLevel: "debug",
      },
      ingress: {
        customDomains: [
          {
            name: "www.my-name.com",
            bindingType: "SniEnabled",
            certificateId:
              "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube/certificates/my-certificate-for-my-name-dot-com",
          },
          {
            name: "www.my-other-name.com",
            bindingType: "SniEnabled",
            certificateId:
              "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube/certificates/my-certificate-for-my-other-name-dot-com",
          },
        ],
        external: true,
        ipSecurityRestrictions: [
          {
            name: "Allow work IP A subnet",
            description: "Allowing all IP's within the subnet below to access containerapp",
            action: "Allow",
            ipAddressRange: "192.168.1.1/32",
          },
          {
            name: "Allow work IP B subnet",
            description: "Allowing all IP's within the subnet below to access containerapp",
            action: "Allow",
            ipAddressRange: "192.168.1.1/8",
          },
        ],
        stickySessions: { affinity: "sticky" },
        targetPort: 3000,
        traffic: [
          {
            label: "production",
            revisionName: "testcontainerApp0-ab1234",
            weight: 100,
          },
        ],
      },
      maxInactiveRevisions: 10,
      runtime: {
        dotnet: { autoConfigureDataProtection: true },
        java: {
          enableMetrics: true,
          javaAgent: {
            enabled: true,
            logging: {
              loggerSettings: [{ level: "debug", logger: "org.springframework.boot" }],
            },
          },
        },
      },
      service: { type: "redis" },
    },
    location: "East US",
    tags: { tag1: "value1", tag2: "value2" },
    template: {
      containers: [
        {
          name: "testcontainerApp0",
          image: "repo/testcontainerApp0:v1",
          probes: [
            {
              type: "Liveness",
              httpGet: {
                path: "/health",
                httpHeaders: [{ name: "Custom-Header", value: "Awesome" }],
                port: 8080,
              },
              initialDelaySeconds: 3,
              periodSeconds: 3,
            },
          ],
        },
      ],
      initContainers: [
        {
          name: "testinitcontainerApp0",
          image: "repo/testcontainerApp0:v4",
          resources: { cpu: 0.2, memory: "100Mi" },
        },
      ],
      scale: {
        cooldownPeriod: 350,
        maxReplicas: 5,
        minReplicas: 1,
        pollingInterval: 35,
        rules: [
          {
            name: "httpscalingrule",
            custom: { type: "http", metadata: { concurrentRequests: "50" } },
          },
        ],
      },
      serviceBinds: [
        {
          name: "service",
          clientType: "dotnet",
          customizedKeys: { desiredKey: "defaultKey" },
          serviceId:
            "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/containerApps/service",
        },
      ],
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.containerApps.beginUpdateAndWait(
    resourceGroupName,
    containerAppName,
    containerAppEnvelope,
  );
  console.log(result);
}
