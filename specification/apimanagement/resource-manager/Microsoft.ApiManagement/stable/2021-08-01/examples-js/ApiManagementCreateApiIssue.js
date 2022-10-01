const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new Issue for an API or updates an existing one.
 *
 * @summary Creates a new Issue for an API or updates an existing one.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateApiIssue.json
 */
async function apiManagementCreateApiIssue() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const apiId = "57d1f7558aa04f15146d9d8a";
  const issueId = "57d2ef278aa04f0ad01d6cdc";
  const parameters = {
    description: "New API issue description",
    createdDate: new Date("2018-02-01T22:21:20.467Z"),
    state: "open",
    title: "New API issue",
    userId:
      "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/1",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiIssue.createOrUpdate(
    resourceGroupName,
    serviceName,
    apiId,
    issueId,
    parameters
  );
  console.log(result);
}

apiManagementCreateApiIssue().catch(console.error);
