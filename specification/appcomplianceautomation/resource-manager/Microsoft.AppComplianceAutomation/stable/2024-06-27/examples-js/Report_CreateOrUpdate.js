const {
  AppComplianceAutomationToolForMicrosoft365,
} = require("@azure/arm-appcomplianceautomation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a new AppComplianceAutomation report or update an exiting AppComplianceAutomation report.
 *
 * @summary Create a new AppComplianceAutomation report or update an exiting AppComplianceAutomation report.
 * x-ms-original-file: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Report_CreateOrUpdate.json
 */
async function reportCreateOrUpdate() {
  const reportName = "testReportName";
  const credential = new DefaultAzureCredential();
  const client = new AppComplianceAutomationToolForMicrosoft365(credential);
  const result = await client.report.beginCreateOrUpdateAndWait(reportName, {
    properties: {
      offerGuid: "00000000-0000-0000-0000-000000000001,00000000-0000-0000-0000-000000000002",
      resources: [
        {
          resourceId:
            "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/SignalR/mySignalRService",
          resourceOrigin: "Azure",
          resourceType: "Microsoft.SignalRService/SignalR",
        },
      ],
      storageInfo: {
        accountName: "testStorageAccount",
        location: "East US",
        resourceGroup: "testResourceGroup",
        subscriptionId: "00000000-0000-0000-0000-000000000000",
      },
      timeZone: "GMT Standard Time",
      triggerTime: new Date("2022-03-04T05:00:00.000Z"),
    },
  });
  console.log(result);
}
