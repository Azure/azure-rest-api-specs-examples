const { LoadTestClient } = require("@azure/arm-loadtestservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update LoadTest resource.
 *
 * @summary Create or update LoadTest resource.
 * x-ms-original-file: specification/loadtestservice/resource-manager/Microsoft.LoadTestService/preview/2021-12-01-preview/examples/LoadTests_CreateOrUpdate.json
 */
async function loadTestsCreateOrUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "dummyrg";
  const loadTestName = "myLoadTest";
  const loadTestResource = {
    description: "This is new load test resource",
    location: "westus",
    tags: { team: "Dev Exp" },
  };
  const credential = new DefaultAzureCredential();
  const client = new LoadTestClient(credential, subscriptionId);
  const result = await client.loadTests.createOrUpdate(
    resourceGroupName,
    loadTestName,
    loadTestResource
  );
  console.log(result);
}

loadTestsCreateOrUpdate().catch(console.error);
