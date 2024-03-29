const { AzureReservationAPI } = require("@azure/arm-reservations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the regions and skus that are available for RI purchase for the specified Azure subscription.
 *
 * @summary Get the regions and skus that are available for RI purchase for the specified Azure subscription.
 * x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-11-01/examples/GetCatalog.json
 */
async function catalog() {
  const subscriptionId = "23bc208b-083f-4901-ae85-4f98c0c3b4b6";
  const reservedResourceType = "VirtualMachines";
  const location = "eastus";
  const options = { reservedResourceType, location };
  const credential = new DefaultAzureCredential();
  const client = new AzureReservationAPI(credential);
  const resArray = new Array();
  for await (let item of client.listCatalog(subscriptionId, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
