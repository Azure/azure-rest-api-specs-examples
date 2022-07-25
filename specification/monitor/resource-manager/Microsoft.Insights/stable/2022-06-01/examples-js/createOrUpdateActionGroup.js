const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a new action group or update an existing one.
 *
 * @summary Create a new action group or update an existing one.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2022-06-01/examples/createOrUpdateActionGroup.json
 */
async function createOrUpdateAnActionGroup() {
  const subscriptionId = "187f412d-1758-44d9-b052-169e2564721d";
  const resourceGroupName = "Default-NotificationRules";
  const actionGroupName = "SampleActionGroup";
  const actionGroup = {
    armRoleReceivers: [
      {
        name: "Sample armRole",
        roleId: "8e3af657-a8ff-443c-a75c-2fe8c4bcb635",
        useCommonAlertSchema: true,
      },
    ],
    automationRunbookReceivers: [
      {
        name: "testRunbook",
        automationAccountId:
          "/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/runbookTest/providers/Microsoft.Automation/automationAccounts/runbooktest",
        isGlobalRunbook: false,
        runbookName: "Sample runbook",
        serviceUri: "<serviceUri>",
        useCommonAlertSchema: true,
        webhookResourceId:
          "/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/runbookTest/providers/Microsoft.Automation/automationAccounts/runbooktest/webhooks/Alert1510184037084",
      },
    ],
    azureAppPushReceivers: [{ name: "Sample azureAppPush", emailAddress: "johndoe@email.com" }],
    azureFunctionReceivers: [
      {
        name: "Sample azureFunction",
        functionAppResourceId:
          "/subscriptions/5def922a-3ed4-49c1-b9fd-05ec533819a3/resourceGroups/aznsTest/providers/Microsoft.Web/sites/testFunctionApp",
        functionName: "HttpTriggerCSharp1",
        httpTriggerUrl: "http://test.me",
        useCommonAlertSchema: true,
      },
    ],
    emailReceivers: [
      {
        name: "John Doe's email",
        emailAddress: "johndoe@email.com",
        useCommonAlertSchema: false,
      },
      {
        name: "Jane Smith's email",
        emailAddress: "janesmith@email.com",
        useCommonAlertSchema: true,
      },
    ],
    enabled: true,
    eventHubReceivers: [
      {
        name: "Sample eventHub",
        eventHubName: "testEventHub",
        eventHubNameSpace: "testEventHubNameSpace",
        subscriptionId: "187f412d-1758-44d9-b052-169e2564721d",
        tenantId: "68a4459a-ccb8-493c-b9da-dd30457d1b84",
      },
    ],
    groupShortName: "sample",
    itsmReceivers: [
      {
        name: "Sample itsm",
        connectionId: "a3b9076c-ce8e-434e-85b4-aff10cb3c8f1",
        region: "westcentralus",
        ticketConfiguration:
          '{"PayloadRevision":0,"WorkItemType":"Incident","UseTemplate":false,"WorkItemData":"{}","CreateOneWIPerCI":false}',
        workspaceId: "5def922a-3ed4-49c1-b9fd-05ec533819a3|55dfd1f8-7e59-4f89-bf56-4c82f5ace23c",
      },
    ],
    location: "Global",
    logicAppReceivers: [
      {
        name: "Sample logicApp",
        callbackUrl:
          "https://prod-27.northcentralus.logic.azure.com/workflows/68e572e818e5457ba898763b7db90877/triggers/manual/paths/invoke/azns/test?api-version=2016-10-01&sp=%2Ftriggers%2Fmanual%2Frun&sv=1.0&sig=Abpsb72UYJxPPvmDo937uzofupO5r_vIeWEx7KVHo7w",
        resourceId:
          "/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/LogicApp/providers/Microsoft.Logic/workflows/testLogicApp",
        useCommonAlertSchema: false,
      },
    ],
    smsReceivers: [
      {
        name: "John Doe's mobile",
        countryCode: "1",
        phoneNumber: "1234567890",
      },
      {
        name: "Jane Smith's mobile",
        countryCode: "1",
        phoneNumber: "0987654321",
      },
    ],
    tags: {},
    voiceReceivers: [{ name: "Sample voice", countryCode: "1", phoneNumber: "1234567890" }],
    webhookReceivers: [
      {
        name: "Sample webhook 1",
        serviceUri: "http://www.example.com/webhook1",
        useCommonAlertSchema: true,
      },
      {
        name: "Sample webhook 2",
        identifierUri: "http://someidentifier/d7811ba3-7996-4a93-99b6-6b2f3f355f8a",
        objectId: "d3bb868c-fe44-452c-aa26-769a6538c808",
        serviceUri: "http://www.example.com/webhook2",
        tenantId: "68a4459a-ccb8-493c-b9da-dd30457d1b84",
        useAadAuth: true,
        useCommonAlertSchema: true,
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.actionGroups.createOrUpdate(
    resourceGroupName,
    actionGroupName,
    actionGroup
  );
  console.log(result);
}

createOrUpdateAnActionGroup().catch(console.error);
