const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or replace an existing cost.
 *
 * @summary Create or replace an existing cost.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Costs_CreateOrUpdate.json
 */
async function costsCreateOrUpdate() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const labName = "{labName}";
  const name = "targetCost";
  const labCost = {
    currencyCode: "USD",
    endDateTime: new Date("2020-12-31T23:59:59Z"),
    startDateTime: new Date("2020-12-01T00:00:00Z"),
    targetCost: {
      costThresholds: [
        {
          displayOnChart: "Disabled",
          percentageThreshold: { thresholdValue: 25 },
          sendNotificationWhenExceeded: "Disabled",
          thresholdId: "00000000-0000-0000-0000-000000000001",
        },
        {
          displayOnChart: "Enabled",
          percentageThreshold: { thresholdValue: 50 },
          sendNotificationWhenExceeded: "Enabled",
          thresholdId: "00000000-0000-0000-0000-000000000002",
        },
        {
          displayOnChart: "Disabled",
          percentageThreshold: { thresholdValue: 75 },
          sendNotificationWhenExceeded: "Disabled",
          thresholdId: "00000000-0000-0000-0000-000000000003",
        },
        {
          displayOnChart: "Disabled",
          percentageThreshold: { thresholdValue: 100 },
          sendNotificationWhenExceeded: "Disabled",
          thresholdId: "00000000-0000-0000-0000-000000000004",
        },
        {
          displayOnChart: "Disabled",
          percentageThreshold: { thresholdValue: 125 },
          sendNotificationWhenExceeded: "Disabled",
          thresholdId: "00000000-0000-0000-0000-000000000005",
        },
      ],
      cycleEndDateTime: new Date("2020-12-31T00:00:00.000Z"),
      cycleStartDateTime: new Date("2020-12-01T00:00:00.000Z"),
      cycleType: "CalendarMonth",
      status: "Enabled",
      target: 100,
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const result = await client.costs.createOrUpdate(resourceGroupName, labName, name, labCost);
  console.log(result);
}

costsCreateOrUpdate().catch(console.error);
