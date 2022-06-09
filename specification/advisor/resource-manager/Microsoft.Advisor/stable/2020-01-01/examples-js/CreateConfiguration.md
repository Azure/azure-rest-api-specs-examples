```javascript
const { AdvisorManagementClient } = require("@azure/arm-advisor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create/Overwrite Azure Advisor configuration and also delete all configurations of contained resource groups.
 *
 * @summary Create/Overwrite Azure Advisor configuration and also delete all configurations of contained resource groups.
 * x-ms-original-file: specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/CreateConfiguration.json
 */
async function putConfigurations() {
  const subscriptionId = "subscriptionId";
  const configurationName = "default";
  const configContract = {
    digests: [
      {
        name: "digestConfigName",
        actionGroupResourceId:
          "/subscriptions/subscriptionId/resourceGroups/resourceGroup/providers/microsoft.insights/actionGroups/actionGroupName",
        categories: [
          "HighAvailability",
          "Security",
          "Performance",
          "Cost",
          "OperationalExcellence",
        ],
        frequency: 30,
        state: "Active",
        language: "en",
      },
    ],
    exclude: true,
    lowCpuThreshold: "5",
  };
  const credential = new DefaultAzureCredential();
  const client = new AdvisorManagementClient(credential, subscriptionId);
  const result = await client.configurations.createInSubscription(
    configurationName,
    configContract
  );
  console.log(result);
}

putConfigurations().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-advisor_3.0.1/sdk/advisor/arm-advisor/README.md) on how to add the SDK to your project and authenticate.
