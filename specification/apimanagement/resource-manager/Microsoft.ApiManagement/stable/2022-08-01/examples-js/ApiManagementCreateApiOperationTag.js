const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Assign tag to the Operation.
 *
 * @summary Assign tag to the Operation.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateApiOperationTag.json
 */
async function apiManagementCreateApiOperationTag() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const apiId = "5931a75ae4bbd512a88c680b";
  const operationId = "5931a75ae4bbd512a88c680a";
  const tagId = "tagId1";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.tag.assignToOperation(
    resourceGroupName,
    serviceName,
    apiId,
    operationId,
    tagId
  );
  console.log(result);
}
