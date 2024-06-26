const { ConfluentManagementClient } = require("@azure/arm-confluent");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Organization Validate proxy resource
 *
 * @summary Organization Validate proxy resource
 * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/Validations_ValidateOrganizationsV2.json
 */
async function validationsValidateOrganizations() {
  const subscriptionId =
    process.env["CONFLUENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["CONFLUENT_RESOURCE_GROUP"] || "myResourceGroup";
  const organizationName = "myOrganization";
  const body = {
    location: "West US",
    offerDetail: {
      id: "string",
      planId: "string",
      planName: "string",
      privateOfferId: "string",
      privateOfferIds: ["string"],
      publisherId: "string",
      termUnit: "string",
    },
    tags: { environment: "Dev" },
    userDetail: {
      aadEmail: "abc@microsoft.com",
      emailAddress: "abc@microsoft.com",
      firstName: "string",
      lastName: "string",
      userPrincipalName: "abc@microsoft.com",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ConfluentManagementClient(credential, subscriptionId);
  const result = await client.validations.validateOrganizationV2(
    resourceGroupName,
    organizationName,
    body,
  );
  console.log(result);
}
