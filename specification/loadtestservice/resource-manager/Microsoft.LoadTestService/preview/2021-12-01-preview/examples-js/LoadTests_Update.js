const { LoadTestClient } = require("@azure/arm-loadtestservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a loadtest resource.
 *
 * @summary Update a loadtest resource.
 * x-ms-original-file: specification/loadtestservice/resource-manager/Microsoft.LoadTestService/preview/2021-12-01-preview/examples/LoadTests_Update.json
 */
async function loadTestsUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "dummyrg";
  const loadTestName = "myLoadTest";
  const loadTestResourcePatchRequestBody = {
    properties: { description: "This is new load test resource" },
    tags: { Division: "LT", Team: "Dev Exp" },
  };
  const credential = new DefaultAzureCredential();
  const client = new LoadTestClient(credential, subscriptionId);
  const result = await client.loadTests.update(
    resourceGroupName,
    loadTestName,
    loadTestResourcePatchRequestBody
  );
  console.log(result);
}

loadTestsUpdate().catch(console.error);
