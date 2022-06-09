```javascript
const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get GitHub Access Token.
 *
 * @summary Get GitHub Access Token.
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Factories_GetGitHubAccessToken.json
 */
async function factoriesGetGitHubAccessToken() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const gitHubAccessTokenRequest = {
    gitHubAccessCode: "some",
    gitHubAccessTokenBaseUrl: "some",
    gitHubClientId: "some",
  };
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.factories.getGitHubAccessToken(
    resourceGroupName,
    factoryName,
    gitHubAccessTokenRequest
  );
  console.log(result);
}

factoriesGetGitHubAccessToken().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-datafactory_10.6.0/sdk/datafactory/arm-datafactory/README.md) on how to add the SDK to your project and authenticate.
