const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get tag associated with the Operation.
 *
 * @summary Get tag associated with the Operation.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetApiOperationTag.json
 */
async function apiManagementGetApiOperationTag() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const apiId = "59d6bb8f1f7fab13dc67ec9b";
  const operationId = "59d6bb8f1f7fab13dc67ec9a";
  const tagId = "59306a29e4bbd510dc24e5f9";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.tag.getByOperation(
    resourceGroupName,
    serviceName,
    apiId,
    operationId,
    tagId
  );
  console.log(result);
}

apiManagementGetApiOperationTag().catch(console.error);
