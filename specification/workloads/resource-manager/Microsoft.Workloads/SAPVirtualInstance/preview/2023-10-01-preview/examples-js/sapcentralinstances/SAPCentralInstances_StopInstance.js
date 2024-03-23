const { WorkloadsClient } = require("@azure/arm-workloadssapvirtualinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Stops the SAP Central Services Instance.
 *
 * @summary Stops the SAP Central Services Instance.
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/examples/sapcentralinstances/SAPCentralInstances_StopInstance.json
 */
async function stopTheSapCentralServicesInstance() {
  const subscriptionId =
    process.env["WORKLOADS_SUBSCRIPTION_ID"] || "8e17e36c-42e9-4cd5-a078-7b44883414e0";
  const resourceGroupName = process.env["WORKLOADS_RESOURCE_GROUP"] || "test-rg";
  const sapVirtualInstanceName = "X00";
  const centralInstanceName = "centralServer";
  const body = { softStopTimeoutSeconds: 1200 };
  const options = { body };
  const credential = new DefaultAzureCredential();
  const client = new WorkloadsClient(credential, subscriptionId);
  const result = await client.sAPCentralInstances.beginStopInstanceAndWait(
    resourceGroupName,
    sapVirtualInstanceName,
    centralInstanceName,
    options,
  );
  console.log(result);
}
