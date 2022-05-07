Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-devtestlabs_4.0.1/sdk/devtestlabs/arm-devtestlabs/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or replace an existing schedule.
 *
 * @summary Create or replace an existing schedule.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Schedules_CreateOrUpdate.json
 */
async function schedulesCreateOrUpdate() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const labName = "{labName}";
  const name = "{scheduleName}";
  const schedule = {
    dailyRecurrence: { time: "{timeOfTheDayTheScheduleWillOccurEveryDay}" },
    hourlyRecurrence: { minute: 30 },
    location: "{location}",
    notificationSettings: {
      emailRecipient: "{email}",
      notificationLocale: "EN",
      status: "{Enabled|Disabled}",
      timeInMinutes: 15,
      webhookUrl: "{webhookUrl}",
    },
    status: "{Enabled|Disabled}",
    tags: { tagName1: "tagValue1" },
    targetResourceId:
      "/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}",
    taskType: "{myLabVmTaskType}",
    timeZoneId: "Pacific Standard Time",
    weeklyRecurrence: {
      time: "{timeOfTheDayTheScheduleWillOccurOnThoseDays}",
      weekdays: ["Monday", "Wednesday", "Friday"],
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const result = await client.schedules.createOrUpdate(resourceGroupName, labName, name, schedule);
  console.log(result);
}

schedulesCreateOrUpdate().catch(console.error);
```
