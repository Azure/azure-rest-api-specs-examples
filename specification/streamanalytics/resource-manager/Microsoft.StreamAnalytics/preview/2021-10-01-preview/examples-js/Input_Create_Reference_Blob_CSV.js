const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates an input or replaces an already existing input under an existing streaming job.
 *
 * @summary Creates an input or replaces an already existing input under an existing streaming job.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Input_Create_Reference_Blob_CSV.json
 */
async function createAReferenceBlobInputWithCsvSerialization() {
  const subscriptionId =
    process.env["STREAMANALYTICS_SUBSCRIPTION_ID"] || "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const resourceGroupName = process.env["STREAMANALYTICS_RESOURCE_GROUP"] || "sjrg8440";
  const jobName = "sj9597";
  const inputName = "input7225";
  const input = {
    properties: {
      type: "Reference",
      datasource: {
        type: "Microsoft.Storage/Blob",
        blobName: "myblobinput",
        container: "state",
        dateFormat: "yyyy/MM/dd",
        deltaPathPattern: "/testBlob/{date}/delta/{time}/",
        deltaSnapshotRefreshRate: "16:14:30",
        fullSnapshotRefreshRate: "16:14:30",
        pathPattern: "{date}/{time}",
        sourcePartitionCount: 16,
        storageAccounts: [{ accountKey: "someAccountKey==", accountName: "someAccountName" }],
        timeFormat: "HH",
      },
      serialization: { type: "Csv", encoding: "UTF8", fieldDelimiter: "," },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.inputs.createOrReplace(resourceGroupName, jobName, inputName, input);
  console.log(result);
}
