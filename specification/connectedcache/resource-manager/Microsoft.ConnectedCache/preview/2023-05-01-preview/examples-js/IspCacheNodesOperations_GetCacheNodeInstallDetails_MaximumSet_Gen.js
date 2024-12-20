const { ConnectedCacheClient } = require("@azure/arm-connectedcache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to this api gets secrets of the ispCacheNode resource install details
 *
 * @summary this api gets secrets of the ispCacheNode resource install details
 * x-ms-original-file: 2023-05-01-preview/IspCacheNodesOperations_GetCacheNodeInstallDetails_MaximumSet_Gen.json
 */
async function ispCacheNodeResourceGetInstallDetailsGeneratedByMaximumSetRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const client = new ConnectedCacheClient(credential, subscriptionId);
  const result = await client.ispCacheNodesOperations.getCacheNodeInstallDetails(
    "rgConnectedCache",
    "MccRPTest1",
    "MCCCachenode1",
  );
  console.log(result);
}
