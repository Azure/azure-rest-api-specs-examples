const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Aborting last running operation on agent pool. We return a 204 no content code here to indicate that the operation has been accepted and an abort will be attempted but is not guaranteed to complete successfully. Please look up the provisioning state of the agent pool to keep track of whether it changes to Canceled. A canceled provisioning state indicates that the abort was successful
 *
 * @summary Aborting last running operation on agent pool. We return a 204 no content code here to indicate that the operation has been accepted and an abort will be attempted but is not guaranteed to complete successfully. Please look up the provisioning state of the agent pool to keep track of whether it changes to Canceled. A canceled provisioning state indicates that the abort was successful
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-08-02-preview/examples/AgentPoolsAbortOperation.json
 */
async function abortOperationOnAgentPool() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "clustername1";
  const agentPoolName = "agentpool1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.agentPools.abortLatestOperation(
    resourceGroupName,
    resourceName,
    agentPoolName
  );
  console.log(result);
}

abortOperationOnAgentPool().catch(console.error);
