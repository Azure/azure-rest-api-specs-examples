const { ConfluentManagementClient } = require("@azure/arm-confluent");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create Organization resource
 *
 * @summary Create Organization resource
 * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/preview/2021-09-01-preview/examples/Organization_Create.json
 */
async function organizationCreate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const organizationName = "myOrganization";
  const body = {
    location: "West US",
    offerDetail: {
      id: "string",
      planId: "string",
      planName: "string",
      publisherId: "string",
      termUnit: "string",
    },
    tags: { environment: "Dev" },
    userDetail: {
      emailAddress: "contoso@microsoft.com",
      firstName: "string",
      lastName: "string",
    },
  };
  const options = { body };
  const credential = new DefaultAzureCredential();
  const client = new ConfluentManagementClient(credential, subscriptionId);
  const result = await client.organization.beginCreateAndWait(
    resourceGroupName,
    organizationName,
    options
  );
  console.log(result);
}

organizationCreate().catch(console.error);
