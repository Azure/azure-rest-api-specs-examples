const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the certificate resource.
 *
 * @summary Get the certificate resource.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-11-01-preview/examples/Certificates_Get.json
 */
async function certificatesGet() {
  const subscriptionId =
    process.env["APPPLATFORM_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APPPLATFORM_RESOURCE_GROUP"] || "myResourceGroup";
  const serviceName = "myservice";
  const certificateName = "mycertificate";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.certificates.get(resourceGroupName, serviceName, certificateName);
  console.log(result);
}
