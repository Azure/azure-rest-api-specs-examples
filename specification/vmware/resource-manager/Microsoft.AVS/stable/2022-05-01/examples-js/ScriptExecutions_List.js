const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List script executions in a private cloud
 *
 * @summary List script executions in a private cloud
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/ScriptExecutions_List.json
 */
async function scriptExecutionsList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "group1";
  const privateCloudName = "{privateCloudName}";
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.scriptExecutions.list(resourceGroupName, privateCloudName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

scriptExecutionsList().catch(console.error);
