const createComputeManagementClient = require("@azure-rest/arm-compute").default,
  { paginate } = require("@azure-rest/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv").config();

/**
 * This sample demonstrates how to This API is deprecated. Use [Resources Skus](https://docs.microsoft.com/rest/api/compute/resourceskus/list)
 *
 * @summary This API is deprecated. Use [Resources Skus](https://docs.microsoft.com/rest/api/compute/resourceskus/list)
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/computeRPCommonExamples/VirtualMachineSizes_List_MaximumSet_Gen.json
 */
async function virtualMachineSizesListMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const client = createComputeManagementClient(credential);
  const subscriptionId = "";
  const location = "-e";
  const options = {
    queryParameters: { "api-version": "2022-08-01" },
  };
  const initialResponse = await client
    .path(
      "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/vmSizes",
      subscriptionId,
      location,
    )
    .get(options);
  const pageData = paginate(client, initialResponse);
  const result = [];
  for await (const item of pageData) {
    result.push(item);
  }
  console.log(result);
}

virtualMachineSizesListMaximumSetGen().catch(console.error);
