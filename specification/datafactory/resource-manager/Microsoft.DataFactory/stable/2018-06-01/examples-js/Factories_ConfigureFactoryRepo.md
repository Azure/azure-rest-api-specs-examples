Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-datafactory_10.4.0/sdk/datafactory/arm-datafactory/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

async function factoriesConfigureFactoryRepo() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
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

factoriesConfigureFactoryRepo().catch(console.error);
```
