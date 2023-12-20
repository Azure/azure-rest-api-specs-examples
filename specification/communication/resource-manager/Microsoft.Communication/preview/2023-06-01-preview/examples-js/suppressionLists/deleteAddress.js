const { CommunicationServiceManagementClient } = require("@azure/arm-communication");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Operation to delete a single address from a suppression list.
 *
 * @summary Operation to delete a single address from a suppression list.
 * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/preview/2023-06-01-preview/examples/suppressionLists/deleteAddress.json
 */
async function deleteASuppressionListAddressResource() {
  const subscriptionId =
    process.env["COMMUNICATION_SUBSCRIPTION_ID"] || "11112222-3333-4444-5555-666677778888";
  const resourceGroupName = process.env["COMMUNICATION_RESOURCE_GROUP"] || "MyResourceGroup";
  const emailServiceName = "MyEmailServiceResource";
  const domainName = "mydomain.com";
  const suppressionListName = "aaaa1111-bbbb-2222-3333-aaaa11112222";
  const addressId = "11112222-3333-4444-5555-999999999999";
  const credential = new DefaultAzureCredential();
  const client = new CommunicationServiceManagementClient(credential, subscriptionId);
  const result = await client.suppressionListAddresses.delete(
    resourceGroupName,
    emailServiceName,
    domainName,
    suppressionListName,
    addressId
  );
  console.log(result);
}
