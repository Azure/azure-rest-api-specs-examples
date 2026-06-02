const { DataBoxManagementClient } = require("@azure/arm-databox");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates the properties of an existing job.
 *
 * @summary updates the properties of an existing job.
 * x-ms-original-file: 2025-07-01/JobsPatch.json
 */
async function jobsPatch() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "YourSubscriptionId";
  const client = new DataBoxManagementClient(credential, subscriptionId);
  const result = await client.jobs.update("YourResourceGroupName", "TestJobName1", {
    details: {
      contactDetails: {
        contactName: "XXXX XXXX",
        emailList: ["xxxx@xxxx.xxx"],
        phone: "0000000000",
        phoneExtension: "",
      },
      shippingAddress: {
        addressType: "Commercial",
        city: "XXXX XXXX",
        companyName: "XXXX XXXX",
        country: "XX",
        postalCode: "00000",
        stateOrProvince: "XX",
        streetAddress1: "XXXX XXXX",
        streetAddress2: "XXXX XXXX",
      },
    },
  });
  console.log(result);
}
