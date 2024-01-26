const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a run of a pipeline.
 *
 * @summary Creates a run of a pipeline.
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Pipelines_CreateRun.json
 */
async function pipelinesCreateRun() {
  const subscriptionId =
    process.env["DATAFACTORY_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = process.env["DATAFACTORY_RESOURCE_GROUP"] || "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const pipelineName = "examplePipeline";
  const referencePipelineRunId = undefined;
  const parameters = {
    outputBlobNameList: ["exampleoutput.csv"],
  };
  const options = {
    referencePipelineRunId,
    parameters,
  };
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.pipelines.createRun(
    resourceGroupName,
    factoryName,
    pipelineName,
    options,
  );
  console.log(result);
}
