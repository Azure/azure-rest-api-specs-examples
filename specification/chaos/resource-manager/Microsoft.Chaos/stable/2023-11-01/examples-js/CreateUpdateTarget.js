const { ChaosManagementClient } = require("@azure/arm-chaos");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a Target resource that extends a tracked regional resource.
 *
 * @summary Create or update a Target resource that extends a tracked regional resource.
 * x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/stable/2023-11-01/examples/CreateUpdateTarget.json
 */
async function createOrUpdateATargetThatExtendsAVirtualMachineResource() {
  const subscriptionId =
    process.env["CHAOS_SUBSCRIPTION_ID"] || "6b052e15-03d3-4f17-b2e1-be7f07588291";
  const resourceGroupName = process.env["CHAOS_RESOURCE_GROUP"] || "exampleRG";
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
