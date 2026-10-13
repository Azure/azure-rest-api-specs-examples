const { IotHubClient } = require("@azure/arm-iothub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update the metadata of an Iot hub. The usual pattern to modify a property is to retrieve the IoT hub metadata and security metadata, and then combine them with the modified values in a new body to update the IoT hub.
 *
 * @summary create or update the metadata of an Iot hub. The usual pattern to modify a property is to retrieve the IoT hub metadata and security metadata, and then combine them with the modified values in a new body to update the IoT hub.
 * x-ms-original-file: 2026-10-01-preview/CreateOrReplace_IotHub.json
 */
async function createOrReplaceIotHub() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "91d12660-3dec-467a-be2a-213b5544ddc0";
  const client = new IotHubClient(credential, subscriptionId);
  const result = await client.iotHubResource.createOrUpdate("myResourceGroup", "testHub", {
    location: "centraluseuap",
    tags: {},
    etag: "AAAAAAFD6M4=",
    properties: {
      ipFilterRules: [],
      networkRuleSets: {
        defaultAction: "Deny",
        applyToBuiltInEventHubEndpoint: true,
        ipRules: [
          { filterName: "rule1", action: "Allow", ipMask: "131.117.159.53" },
          { filterName: "rule2", action: "Allow", ipMask: "157.55.59.128/25" },
        ],
      },
      eventHubEndpoints: { events: { retentionTimeInDays: 1, partitionCount: 2 } },
      routing: {
        endpoints: {
          serviceBusQueues: [],
          serviceBusTopics: [],
          eventHubs: [],
          storageContainers: [],
        },
        routes: [
          {
            name: "Routeid",
            source: "DeviceMessages",
            condition: "true",
            dataSchema:
              "aio-sr://aiosaalkopkedev/62a24af1d7db61cd44b2ad6b6c3f4ab7312be447f89ff3401d18357d0d05ce3a:1",
            endpointNames: ["events"],
            isEnabled: true,
          },
        ],
        fallbackRoute: {
          name: "$fallback",
          source: "DeviceMessages",
          condition: "true",
          endpointNames: ["events"],
          isEnabled: true,
        },
      },
      storageEndpoints: {
        $default: { sasTtlAsIso8601: "PT1H", connectionString: "", containerName: "" },
      },
      messagingEndpoints: {
        fileNotifications: {
          lockDurationAsIso8601: "PT1M",
          ttlAsIso8601: "PT1H",
          maxDeliveryCount: 10,
        },
      },
      enableFileUploadNotifications: false,
      cloudToDevice: {
        maxDeliveryCount: 10,
        defaultTtlAsIso8601: "PT1H",
        feedback: { lockDurationAsIso8601: "PT1M", ttlAsIso8601: "PT1H", maxDeliveryCount: 10 },
      },
      features: "None",
      minTlsVersion: "1.2",
      enableDataResidency: true,
      rootCertificate: { enableRootCertificateV2: true },
      ipVersion: "ipv4ipv6",
    },
    sku: { name: "S1", capacity: 1 },
  });
  console.log(result);
}
