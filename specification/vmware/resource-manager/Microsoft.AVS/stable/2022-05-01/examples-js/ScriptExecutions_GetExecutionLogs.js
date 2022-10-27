const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Return the logs for a script execution resource
 *
 * @summary Return the logs for a script execution resource
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/ScriptExecutions_GetExecutionLogs.json
 */
async function scriptExecutionsGetExecutionLogs() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "group1";
  const privateCloudName = "cloud1";
  const scriptExecutionName = "addSsoServer";
  const scriptOutputStreamType = ["Information", "Warnings", "Errors", "Output"];
  const options = {
    scriptOutputStreamType,
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.scriptExecutions.getExecutionLogs(
    resourceGroupName,
    privateCloudName,
    scriptExecutionName,
    options
  );
  console.log(result);
}

scriptExecutionsGetExecutionLogs().catch(console.error);
