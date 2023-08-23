const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or Update a Certificate.
 *
 * @summary Create or Update a Certificate.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/Certificate_CreateOrUpdate.json
 */
async function createOrUpdateCertificate() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPCONTAINERS_RESOURCE_GROUP"] || "examplerg";
  const environmentName = "testcontainerenv";
  const certificateName = "certificate-firendly-name";
  const certificateEnvelope = {
    location: "East US",
    properties: {
      password: "private key password",
      value: Buffer.from("Y2VydA=="),
    },
  };
  const options = {
    certificateEnvelope,
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.certificates.createOrUpdate(
    resourceGroupName,
    environmentName,
    certificateName,
    options
  );
  console.log(result);
}
