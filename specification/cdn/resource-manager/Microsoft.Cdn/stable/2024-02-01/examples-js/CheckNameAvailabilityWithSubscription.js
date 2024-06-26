const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Check the availability of a resource name. This is needed for resources where name is globally unique, such as a CDN endpoint.
 *
 * @summary Check the availability of a resource name. This is needed for resources where name is globally unique, such as a CDN endpoint.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/CheckNameAvailabilityWithSubscription.json
 */
async function checkNameAvailabilityWithSubscription() {
  const subscriptionId = process.env["CDN_SUBSCRIPTION_ID"] || "subid";
  const checkNameAvailabilityInput = {
    name: "sampleName",
    type: "Microsoft.Cdn/Profiles/Endpoints",
  };
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.checkNameAvailabilityWithSubscription(checkNameAvailabilityInput);
  console.log(result);
}
