const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or replace an existing schedule.
 *
 * @summary Create or replace an existing schedule.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachineSchedules_CreateOrUpdate.json
 */
async function virtualMachineSchedulesCreateOrUpdate() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const labName = "{labName}";
  const virtualMachineName = "{vmName}";
  const name = "LabVmsShutdown";
  const schedule = {
    dailyRecurrence: { time: "1900" },
    hourlyRecurrence: { minute: 30 },
    location: "{location}",
    notificationSettings: {
      emailRecipient: "{email}",
      notificationLocale: "EN",
      status: "Enabled",
      timeInMinutes: 30,
      webhookUrl: "{webhookUrl}",
    },
    status: "Enabled",
    tags: { tagName1: "tagValue1" },
    targetResourceId:
      "/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/virtualMachines/{vmName}",
    taskType: "LabVmsShutdownTask",
    timeZoneId: "Pacific Standard Time",
    weeklyRecurrence: {
      time: "1700",
      weekdays: ["Friday", "Saturday", "Sunday"],
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const result = await client.virtualMachineSchedules.createOrUpdate(
    resourceGroupName,
    labName,
    virtualMachineName,
    name,
    schedule
  );
  console.log(result);
}

virtualMachineSchedulesCreateOrUpdate().catch(console.error);
