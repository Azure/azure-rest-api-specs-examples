const { DataBoxEdgeManagementClient } = require("@azure/arm-databoxedge");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an order.
 *
 * @summary Creates or updates an order.
 * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/OrderPut.json
 */
async function orderPut() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const deviceName = "testedgedevice";
  const resourceGroupName = "GroupForEdgeAutomation";
  const order = {
    contactInformation: {
      companyName: "Microsoft",
      contactPerson: "John Mcclane",
      emailList: ["john@microsoft.com"],
      phone: "(800) 426-9400",
    },
    shippingAddress: {
      addressLine1: "Microsoft Corporation",
      addressLine2: "One Microsoft Way",
      addressLine3: "Redmond",
      city: "WA",
      country: "USA",
      postalCode: "98052",
      state: "WA",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new DataBoxEdgeManagementClient(credential, subscriptionId);
  const result = await client.orders.beginCreateOrUpdateAndWait(
    deviceName,
    resourceGroupName,
    order
  );
  console.log(result);
}

orderPut().catch(console.error);
