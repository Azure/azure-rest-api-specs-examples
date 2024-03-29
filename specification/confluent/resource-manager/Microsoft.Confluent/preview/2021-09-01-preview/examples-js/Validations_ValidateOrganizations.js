const { ConfluentManagementClient } = require("@azure/arm-confluent");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Organization Validate proxy resource
 *
 * @summary Organization Validate proxy resource
 * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/preview/2021-09-01-preview/examples/Validations_ValidateOrganizations.json
 */
async function validationsValidateOrganizations() {
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
      emailAddress: "abc@microsoft.com",
      firstName: "string",
      lastName: "string",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ConfluentManagementClient(credential, subscriptionId);
  const result = await client.validations.validateOrganization(
    resourceGroupName,
    organizationName,
    body
  );
  console.log(result);
}

validationsValidateOrganizations().catch(console.error);
