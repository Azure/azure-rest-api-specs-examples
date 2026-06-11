const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates new or updates existing specified Schema of the API Management service instance.
 *
 * @summary creates new or updates existing specified Schema of the API Management service instance.
 * x-ms-original-file: 2025-09-01-preview/ApiManagementCreateGlobalSchema2.json
 */
async function apiManagementCreateSchema2() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.globalSchema.createOrUpdate("rg1", "apimService1", "schema1", {
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
  });
  console.log(result);
}
