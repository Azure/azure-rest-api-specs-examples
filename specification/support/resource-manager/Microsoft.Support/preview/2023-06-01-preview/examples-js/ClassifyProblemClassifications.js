const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Classify the right problem classifications (categories) available for a specific Azure service.
 *
 * @summary Classify the right problem classifications (categories) available for a specific Azure service.
 * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/ClassifyProblemClassifications.json
 */
async function classifyListOfProblemClassificationsForASpecifiedAzureService() {
  const problemServiceName = "serviceId1";
  const problemClassificationsClassificationInput = {
    issueSummary: "Can not connect to Windows VM",
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential);
  const result = await client.problemClassificationsNoSubscription.classifyProblems(
    problemServiceName,
    problemClassificationsClassificationInput,
  );
  console.log(result);
}
