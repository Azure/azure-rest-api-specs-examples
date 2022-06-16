const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete the certificate resource.
 *
 * @summary Delete the certificate resource.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/Certificates_Delete.json
 */
async function certificatesDelete() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const certificateName = "mycertificate";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.certificates.beginDeleteAndWait(
    resourceGroupName,
    serviceName,
    certificateName
  );
  console.log(result);
}

certificatesDelete().catch(console.error);
