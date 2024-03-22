const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This operation fetches ARM resource id of support resource type.
 *
 * @summary This operation fetches ARM resource id of support resource type.
 * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/LookUpResourceId.json
 */
async function lookUpResourceIdOfSupportResourceType() {
  const lookUpResourceIdRequest = {
    type: "Microsoft.Support/supportTickets",
    identifier: "1234668596",
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential);
  const result = await client.lookUpResourceId.post(lookUpResourceIdRequest);
  console.log(result);
}
