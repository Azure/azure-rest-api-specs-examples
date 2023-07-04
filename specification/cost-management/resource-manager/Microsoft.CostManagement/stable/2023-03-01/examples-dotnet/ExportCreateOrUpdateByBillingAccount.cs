using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CostManagement;
using Azure.ResourceManager.CostManagement.Models;

// Generated from example definition: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2023-03-01/examples/ExportCreateOrUpdateByBillingAccount.json
// this example is just showing the usage of "Exports_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CostManagementExportResource created on azure
// for more information of creating CostManagementExportResource, please refer to the document of CostManagementExportResource
string scope = "providers/Microsoft.Billing/billingAccounts/123456";
string exportName = "TestExport";
ResourceIdentifier costManagementExportResourceId = CostManagementExportResource.CreateResourceIdentifier(scope, exportName);
CostManagementExportResource costManagementExport = client.GetCostManagementExportResource(costManagementExportResourceId);

// invoke the operation
CostManagementExportData data = new CostManagementExportData()
{
    Format = ExportFormatType.Csv,
    DeliveryInfoDestination = new ExportDeliveryDestination("exports")
    {
        ResourceId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Storage/storageAccounts/ccmeastusdiag182"),
        RootFolderPath = "ad-hoc",
    },
    Definition = new ExportDefinition(ExportType.ActualCost, TimeframeType.MonthToDate)
    {
        DataSet = new ExportDataset()
        {
            Granularity = GranularityType.Daily,
            Columns =
            {
            "Date","MeterId","ResourceId","ResourceLocation","Quantity"
            },
        },
    },
    Schedule = new ExportSchedule()
    {
        Status = ExportScheduleStatusType.Active,
        Recurrence = ExportScheduleRecurrenceType.Weekly,
        RecurrencePeriod = new ExportRecurrencePeriod(DateTimeOffset.Parse("2020-06-01T00:00:00Z"))
        {
            To = DateTimeOffset.Parse("2020-10-31T00:00:00Z"),
        },
    },
};
ArmOperation<CostManagementExportResource> lro = await costManagementExport.UpdateAsync(WaitUntil.Completed, data);
CostManagementExportResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CostManagementExportData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
