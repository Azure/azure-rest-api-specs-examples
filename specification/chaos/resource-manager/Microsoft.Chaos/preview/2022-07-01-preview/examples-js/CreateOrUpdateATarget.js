const { ChaosManagementClient } = require("@azure/arm-chaos");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a Target resource that extends a tracked regional resource.
 *
 * @summary Create or update a Target resource that extends a tracked regional resource.
 * x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/preview/2022-07-01-preview/examples/CreateOrUpdateATarget.json
 */
async function createOrUpdateATargetThatExtendsAVirtualMachineResource() {
  const subscriptionId = "6b052e15-03d3-4f17-b2e1-be7f07588291";
  const resourceGroupName = "exampleRG";
  const parentProviderNamespace = "Microsoft.Compute";
  const parentResourceType = "virtualMachines";
  const parentResourceName = "exampleVM";
  const targetName = "Microsoft-Agent";
  const target = {
    properties: {
      identities: [{ type: "CertificateSubjectIssuer", subject: "CN=example.subject" }],
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ChaosManagementClient(credential, subscriptionId);
  const result = await client.targets.createOrUpdate(
    resourceGroupName,
    parentProviderNamespace,
    parentResourceType,
    parentResourceName,
    targetName,
    target
  );
  console.log(result);
}

createOrUpdateATargetThatExtendsAVirtualMachineResource().catch(console.error);
