const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List Online Endpoints.
 *
 * @summary List Online Endpoints.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/OnlineEndpoint/list.json
 */
async function listOnlineEndpoint() {
  const subscriptionId =
    process.env["MACHINELEARNING_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["MACHINELEARNING_RESOURCE_GROUP"] || "test-rg";
  const workspaceName = "my-aml-workspace";
  const name = "string";
  const count = 1;
  const computeType = "Managed";
  const tags = "string";
  const properties = "string";
  const orderBy = "CreatedAtDesc";
  const options = {
    name,
    count,
    computeType,
    tags,
    properties,
    orderBy,
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.onlineEndpoints.list(resourceGroupName, workspaceName, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
