const { LoadTestClient } = require("@azure/arm-loadtesting");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the endpoints that agents may call as part of load testing.
 *
 * @summary Lists the endpoints that agents may call as part of load testing.
 * x-ms-original-file: specification/loadtestservice/resource-manager/Microsoft.LoadTestService/stable/2022-12-01/examples/LoadTests_ListOutboundNetworkDependenciesEndpoints.json
 */
async function listOutboundNetworkDependencies() {
  const subscriptionId = "subid";
  const resourceGroupName = "default-azureloadtest-japaneast";
  const loadTestName = "sampleloadtest";
  const credential = new DefaultAzureCredential();
  const client = new LoadTestClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.loadTests.listOutboundNetworkDependenciesEndpoints(
    resourceGroupName,
    loadTestName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listOutboundNetworkDependencies().catch(console.error);
