const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates new or updates existing specified Schema of the API Management service instance.
 *
 * @summary Creates new or updates existing specified Schema of the API Management service instance.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateGlobalSchema2.json
 */
async function apiManagementCreateSchema2() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const schemaId = "schema1";
  const parameters = {
    description: "sample schema description",
    document: {
      type: "object",
      $id: "https://example.com/person.schema.json",
      $schema: "https://json-schema.org/draft/2020-12/schema",
      properties: {
        age: {
          type: "integer",
          description: "Age in years which must be equal to or greater than zero.",
          minimum: 0,
        },
        firstName: { type: "string", description: "The person's first name." },
        lastName: { type: "string", description: "The person's last name." },
      },
      title: "Person",
    },
    schemaType: "json",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.globalSchema.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serviceName,
    schemaId,
    parameters
  );
  console.log(result);
}
