const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified service.
 *
 * @summary Deletes the specified service.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/ServiceDelete.json
 */
async function deleteService() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const mobileNetworkName = "testMobileNetwork";
  const serviceName = "TestService";
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.services.beginDeleteAndWait(
    resourceGroupName,
    mobileNetworkName,
    serviceName
  );
  console.log(result);
}

deleteService().catch(console.error);
