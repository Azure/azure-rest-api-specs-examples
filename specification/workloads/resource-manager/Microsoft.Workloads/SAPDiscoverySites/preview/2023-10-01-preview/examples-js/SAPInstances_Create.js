const { WorkloadsClient } = require("@azure/arm-migrationdiscoverysap");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates the SAP Instance resource. <br><br>This will be used by service only. PUT operation on this resource by end user will return a Bad Request error.
 *
 * @summary Creates the SAP Instance resource. <br><br>This will be used by service only. PUT operation on this resource by end user will return a Bad Request error.
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/SAPDiscoverySites/preview/2023-10-01-preview/examples/SAPInstances_Create.json
 */
async function createsTheSapInstanceResource() {
  const subscriptionId =
    process.env["MIGRATIONDISCOVERY_SUBSCRIPTION_ID"] || "6d875e77-e412-4d7d-9af4-8895278b4443";
  const resourceGroupName = process.env["MIGRATIONDISCOVERY_RESOURCE_GROUP"] || "test-rg";
  const sapDiscoverySiteName = "SampleSite";
  const sapInstanceName = "MPP_MPP";
  const resource = {
    location: "eastus",
    properties: {},
    tags: { property1: "value1", property2: "value2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new WorkloadsClient(credential, subscriptionId);
  const result = await client.sapInstances.beginCreateAndWait(
    resourceGroupName,
    sapDiscoverySiteName,
    sapInstanceName,
    resource,
  );
  console.log(result);
}
