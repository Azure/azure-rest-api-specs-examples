const { AgriFoodMgmtClient } = require("@azure/arm-agrifood");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get operationResults for a FarmBeats resource.
 *
 * @summary Get operationResults for a FarmBeats resource.
 * x-ms-original-file: specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/FarmBeatsModels_GetOperationResult.json
 */
async function farmBeatsModelsGetOperationResult() {
  const subscriptionId = "11111111-2222-3333-4444-555555555555";
  const resourceGroupName = "examples-rg";
  const farmBeatsResourceName = "examples-farmBeatsResourceName";
  const operationResultsId = "resource-provisioning-id-farmBeatsResourceName";
  const credential = new DefaultAzureCredential();
  const client = new AgriFoodMgmtClient(credential, subscriptionId);
  const result = await client.farmBeatsModels.getOperationResult(
    resourceGroupName,
    farmBeatsResourceName,
    operationResultsId
  );
  console.log(result);
}

farmBeatsModelsGetOperationResult().catch(console.error);
