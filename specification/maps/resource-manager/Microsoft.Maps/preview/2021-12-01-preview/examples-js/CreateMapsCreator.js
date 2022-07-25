const { AzureMapsManagementClient } = require("@azure/arm-maps");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a Maps Creator resource. Creator resource will manage Azure resources required to populate a custom set of mapping data. It requires an account to exist before it can be created.
 *
 * @summary Create or update a Maps Creator resource. Creator resource will manage Azure resources required to populate a custom set of mapping data. It requires an account to exist before it can be created.
 * x-ms-original-file: specification/maps/resource-manager/Microsoft.Maps/preview/2021-12-01-preview/examples/CreateMapsCreator.json
 */
async function createCreatorResource() {
  const subscriptionId = "21a9967a-e8a9-4656-a70b-96ff1c4d05a0";
  const resourceGroupName = "myResourceGroup";
  const accountName = "myMapsAccount";
  const creatorName = "myCreator";
  const creatorResource = {
    location: "eastus2",
    properties: { storageUnits: 5 },
    tags: { test: "true" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMapsManagementClient(credential, subscriptionId);
  const result = await client.creators.createOrUpdate(
    resourceGroupName,
    accountName,
    creatorName,
    creatorResource
  );
  console.log(result);
}

createCreatorResource().catch(console.error);
