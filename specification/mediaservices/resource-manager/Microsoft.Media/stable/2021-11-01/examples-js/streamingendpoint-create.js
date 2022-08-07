const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a streaming endpoint.
 *
 * @summary Creates a streaming endpoint.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/streamingendpoint-create.json
 */
async function createAStreamingEndpoint() {
  const subscriptionId = "0a6ec948-5a62-437d-b9df-934dc7c1b722";
  const resourceGroupName = "mediaresources";
  const accountName = "slitestmedia10";
  const streamingEndpointName = "myStreamingEndpoint1";
  const parameters = {
    description: "test event 1",
    accessControl: {
      akamai: {
        akamaiSignatureHeaderAuthenticationKeyList: [
          {
            base64Key: "dGVzdGlkMQ==",
            expiration: new Date("2029-12-31T16:00:00-08:00"),
            identifier: "id1",
          },
          {
            base64Key: "dGVzdGlkMQ==",
            expiration: new Date("2030-12-31T16:00:00-08:00"),
            identifier: "id2",
          },
        ],
      },
      ip: { allow: [{ name: "AllowedIp", address: "192.168.1.1" }] },
    },
    availabilitySetName: "availableset",
    cdnEnabled: false,
    location: "West US",
    scaleUnits: 1,
    tags: { tag1: "value1", tag2: "value2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.streamingEndpoints.beginCreateAndWait(
    resourceGroupName,
    accountName,
    streamingEndpointName,
    parameters
  );
  console.log(result);
}

createAStreamingEndpoint().catch(console.error);
