const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

async function submitNewCommand() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "clustername1";
  const requestPayload = {
    clusterToken: "",
    command: "kubectl apply -f ns.yaml",
    context: "",
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.managedClusters.beginRunCommandAndWait(
    resourceGroupName,
    resourceName,
    requestPayload
  );
  console.log(result);
}

submitNewCommand().catch(console.error);
