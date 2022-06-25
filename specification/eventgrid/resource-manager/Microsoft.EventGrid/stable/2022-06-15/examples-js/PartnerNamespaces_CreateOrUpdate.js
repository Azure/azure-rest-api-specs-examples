const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Asynchronously creates a new partner namespace with the specified parameters.
 *
 * @summary Asynchronously creates a new partner namespace with the specified parameters.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/PartnerNamespaces_CreateOrUpdate.json
 */
async function partnerNamespacesCreateOrUpdate() {
  const subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
  const resourceGroupName = "examplerg";
  const partnerNamespaceName = "examplePartnerNamespaceName1";
  const partnerNamespaceInfo = {
    location: "westus",
    partnerRegistrationFullyQualifiedId:
      "/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.EventGrid/partnerRegistrations/ContosoCorpAccount1",
    tags: { tag1: "value1", tag2: "value2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const result = await client.partnerNamespaces.beginCreateOrUpdateAndWait(
    resourceGroupName,
    partnerNamespaceName,
    partnerNamespaceInfo
  );
  console.log(result);
}

partnerNamespacesCreateOrUpdate().catch(console.error);
