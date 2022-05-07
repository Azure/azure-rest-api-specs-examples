Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-iothub_6.1.1/sdk/iothub/arm-iothub/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { IotHubClient } = require("@azure/arm-iothub");
const { DefaultAzureCredential } = require("@azure/identity");

async function iotHubResourceCreateOrUpdate() {
  const subscriptionId = "91d12660-3dec-467a-be2a-213b5544ddc0";
  const resourceGroupName = "myResourceGroup";
  const resourceName = "testHub";
  const iotHubDescription = {
    etag: "AAAAAAFD6M4=",
    location: "centraluseuap",
    properties: {
      cloudToDevice: {
        defaultTtlAsIso8601: "PT1H",
        feedback: {
          lockDurationAsIso8601: "PT1M",
          maxDeliveryCount: 10,
          ttlAsIso8601: "PT1H",
        },
        maxDeliveryCount: 10,
      },
      enableDataResidency: true,
      enableFileUploadNotifications: false,
      eventHubEndpoints: {
        events: { partitionCount: 2, retentionTimeInDays: 1 },
      },
      features: "None",
      ipFilterRules: [],
      messagingEndpoints: {
        fileNotifications: {
          lockDurationAsIso8601: "PT1M",
          maxDeliveryCount: 10,
          ttlAsIso8601: "PT1H",
        },
      },
      minTlsVersion: "1.2",
      networkRuleSets: {
        applyToBuiltInEventHubEndpoint: true,
        defaultAction: "Deny",
        ipRules: [
          { action: "Allow", filterName: "rule1", ipMask: "131.117.159.53" },
          { action: "Allow", filterName: "rule2", ipMask: "157.55.59.128/25" },
        ],
      },
      routing: {
        endpoints: {
          eventHubs: [],
          serviceBusQueues: [],
          serviceBusTopics: [],
          storageContainers: [],
        },
        fallbackRoute: {
          name: "$fallback",
          condition: "true",
          endpointNames: ["events"],
          isEnabled: true,
          source: "DeviceMessages",
        },
        routes: [],
      },
      storageEndpoints: {
        $default: {
          connectionString: "",
          containerName: "",
          sasTtlAsIso8601: "PT1H",
        },
      },
    },
    sku: { name: "S1", capacity: 1 },
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new IotHubClient(credential, subscriptionId);
  const result = await client.iotHubResource.beginCreateOrUpdateAndWait(
    resourceGroupName,
    resourceName,
    iotHubDescription
  );
  console.log(result);
}

iotHubResourceCreateOrUpdate().catch(console.error);
```
