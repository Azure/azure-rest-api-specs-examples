const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a factory's repo information.
 *
 * @summary Updates a factory's repo information.
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Factories_ConfigureFactoryRepo.json
 */
async function factoriesConfigureFactoryRepo() {
  const subscriptionId =
    process.env["DATAFACTORY_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-12345678abc";
  const locationId = "East US";
  const factoryRepoUpdate = {
    factoryResourceId:
      "/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName",
    repoConfiguration: {
      type: "FactoryVSTSConfiguration",
      accountName: "ADF",
      collaborationBranch: "master",
      lastCommitId: "",
      projectName: "project",
      repositoryName: "repo",
      rootFolder: "/",
      tenantId: "",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.factories.configureFactoryRepo(locationId, factoryRepoUpdate);
  console.log(result);
}
