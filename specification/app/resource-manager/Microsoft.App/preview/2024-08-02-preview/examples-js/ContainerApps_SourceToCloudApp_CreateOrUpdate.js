const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a Container App.
 *
 * @summary Create or update a Container App.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/ContainerApps_SourceToCloudApp_CreateOrUpdate.json
 */
async function createOrUpdateSourceToCloudApp() {
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
        additionalPortMappings: [
          { external: true, targetPort: 1234 },
          { exposedPort: 3456, external: false, targetPort: 2345 },
        ],
        clientCertificateMode: "accept",
        corsPolicy: {
          allowCredentials: true,
          allowedHeaders: ["HEADER1", "HEADER2"],
          allowedMethods: ["GET", "POST"],
          allowedOrigins: ["https://a.test.com", "https://b.test.com"],
          exposeHeaders: ["HEADER3", "HEADER4"],
          maxAge: 1234,
        },
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
      service: { type: "redis" },
    },
    environmentId:
      "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube",
    location: "East US",
    patchingConfiguration: { patchingMode: "Automatic" },
    template: {
      containers: [
        {
          name: "testcontainerApp0",
          image: "",
          imageType: "CloudBuild",
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
          volumeMounts: [
            {
              mountPath: "/mnt/path1",
              subPath: "subPath1",
              volumeName: "azurefile",
            },
            {
              mountPath: "/mnt/path2",
              subPath: "subPath2",
              volumeName: "nfsazurefile",
            },
          ],
        },
      ],
      initContainers: [
        {
          name: "testinitcontainerApp0",
          args: ["-c", "while true; do echo hello; sleep 10;done"],
          command: ["/bin/sh"],
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
          name: "redisService",
          clientType: "dotnet",
          customizedKeys: { desiredKey: "defaultKey" },
          serviceId:
            "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/containerApps/redisService",
        },
      ],
      volumes: [
        { name: "azurefile", storageName: "storage", storageType: "AzureFile" },
        {
          name: "nfsazurefile",
          storageName: "nfsStorage",
          storageType: "NfsAzureFile",
        },
      ],
    },
    workloadProfileName: "My-GP-01",
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.containerApps.beginCreateOrUpdateAndWait(
    resourceGroupName,
    containerAppName,
    containerAppEnvelope,
  );
  console.log(result);
}
