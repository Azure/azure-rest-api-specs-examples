const { LogicManagementClient } = require("@azure/arm-logic");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an integration account certificate.
 *
 * @summary Creates or updates an integration account certificate.
 * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountCertificates_CreateOrUpdate.json
 */
async function createOrUpdateACertificate() {
  const subscriptionId =
    process.env["LOGIC_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["LOGIC_RESOURCE_GROUP"] || "testResourceGroup";
  const integrationAccountName = "testIntegrationAccount";
  const certificateName = "testCertificate";
  const certificate = {
    key: {
      keyName: "<keyName>",
      keyVault: {
        id: "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/testResourceGroup/providers/microsoft.keyvault/vaults/<keyVaultName>",
      },
      keyVersion: "87d9764197604449b9b8eb7bd8710868",
    },
    location: "brazilsouth",
    publicCertificate: "<publicCertificateValue>",
  };
  const credential = new DefaultAzureCredential();
  const client = new LogicManagementClient(credential, subscriptionId);
  const result = await client.integrationAccountCertificates.createOrUpdate(
    resourceGroupName,
    integrationAccountName,
    certificateName,
    certificate
  );
  console.log(result);
}
