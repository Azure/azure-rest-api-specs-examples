const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified Managed Certificate.
 *
 * @summary Deletes the specified Managed Certificate.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/ManagedCertificate_Delete.json
 */
async function deleteCertificate() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPCONTAINERS_RESOURCE_GROUP"] || "examplerg";
  const environmentName = "testcontainerenv";
  const managedCertificateName = "certificate-firendly-name";
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.managedCertificates.delete(
    resourceGroupName,
    environmentName,
    managedCertificateName
  );
  console.log(result);
}
