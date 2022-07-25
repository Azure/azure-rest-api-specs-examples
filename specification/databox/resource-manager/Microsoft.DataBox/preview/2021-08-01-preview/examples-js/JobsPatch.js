const { DataBoxManagementClient } = require("@azure/arm-databox");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the properties of an existing job.
 *
 * @summary Updates the properties of an existing job.
 * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/JobsPatch.json
 */
async function jobsPatch() {
  const subscriptionId = "fa68082f-8ff7-4a25-95c7-ce9da541242f";
  const resourceGroupName = "SdkRg5154";
  const jobName = "SdkJob952";
  const jobResourceUpdateParameter = {
    details: {
      contactDetails: {
        contactName: "Update Job",
        emailList: ["testing@microsoft.com"],
        phone: "1234567890",
        phoneExtension: "1234",
      },
      shippingAddress: {
        addressType: "Commercial",
        city: "San Francisco",
        companyName: "Microsoft",
        country: "US",
        postalCode: "94107",
        stateOrProvince: "CA",
        streetAddress1: "16 TOWNSEND ST",
        streetAddress2: "Unit 1",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new DataBoxManagementClient(credential, subscriptionId);
  const result = await client.jobs.beginUpdateAndWait(
    resourceGroupName,
    jobName,
    jobResourceUpdateParameter
  );
  console.log(result);
}

jobsPatch().catch(console.error);
