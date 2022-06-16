const { DomainServicesResourceProvider } = require("@azure/arm-domainservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The Update OuContainer operation can be used to update the existing OuContainers.
 *
 * @summary The Update OuContainer operation can be used to update the existing OuContainers.
 * x-ms-original-file: specification/domainservices/resource-manager/Microsoft.AAD/stable/2021-05-01/examples/UpdateOuContainer.json
 */
async function updateDomainService() {
  const subscriptionId = "1639790a-76a2-4ac4-98d9-8562f5dfcb4d";
  const resourceGroupName = "OuContainerResourceGroup";
  const domainServiceName = "OuContainer.com";
  const ouContainerName = "OuContainer1";
  const containerAccount = {
    accountName: "AccountName1",
    password: "<password>",
    spn: "Spn1",
  };
  const credential = new DefaultAzureCredential();
  const client = new DomainServicesResourceProvider(credential, subscriptionId);
  const result = await client.ouContainerOperationGrp.beginUpdateAndWait(
    resourceGroupName,
    domainServiceName,
    ouContainerName,
    containerAccount
  );
  console.log(result);
}

updateDomainService().catch(console.error);
