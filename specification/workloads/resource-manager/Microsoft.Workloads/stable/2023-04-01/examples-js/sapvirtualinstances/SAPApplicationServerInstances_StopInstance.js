const { WorkloadsClient } = require("@azure/arm-workloads");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Stops the SAP Application Server Instance.
 *
 * @summary Stops the SAP Application Server Instance.
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/SAPApplicationServerInstances_StopInstance.json
 */
async function stopTheSapApplicationServerInstance() {
  const subscriptionId =
    process.env["WORKLOADS_SUBSCRIPTION_ID"] || "8e17e36c-42e9-4cd5-a078-7b44883414e0";
  const resourceGroupName = process.env["WORKLOADS_RESOURCE_GROUP"] || "test-rg";
  const sapVirtualInstanceName = "X00";
  const applicationInstanceName = "app01";
  const body = { softStopTimeoutSeconds: 0 };
  const options = {
    body,
  };
  const credential = new DefaultAzureCredential();
  const client = new WorkloadsClient(credential, subscriptionId);
  const result = await client.sAPApplicationServerInstances.beginStopInstanceAndWait(
    resourceGroupName,
    sapVirtualInstanceName,
    applicationInstanceName,
    options
  );
  console.log(result);
}
