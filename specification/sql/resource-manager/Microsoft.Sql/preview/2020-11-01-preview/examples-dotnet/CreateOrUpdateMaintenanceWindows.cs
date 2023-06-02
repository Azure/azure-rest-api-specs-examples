using System;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;
using Azure.ResourceManager.Sql.Models;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CreateOrUpdateMaintenanceWindows.json
// this example is just showing the usage of "MaintenanceWindows_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MaintenanceWindowsResource created on azure
// for more information of creating MaintenanceWindowsResource, please refer to the document of MaintenanceWindowsResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "Default-SQL-SouthEastAsia";
string serverName = "testsvr";
string databaseName = "testdwdb";
ResourceIdentifier maintenanceWindowsResourceId = MaintenanceWindowsResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, databaseName);
MaintenanceWindowsResource maintenanceWindows = client.GetMaintenanceWindowsResource(maintenanceWindowsResourceId);

// invoke the operation
string maintenanceWindowName = "current";
MaintenanceWindowsData data = new MaintenanceWindowsData()
{
    TimeRanges =
    {
    new MaintenanceWindowTimeRange()
    {
    DayOfWeek = SqlDayOfWeek.Saturday,
    StartTime = "00:00:00",
    Duration = XmlConvert.ToTimeSpan("PT60M"),
    }
    },
};
await maintenanceWindows.CreateOrUpdateAsync(WaitUntil.Completed, maintenanceWindowName, data);

Console.WriteLine($"Succeeded");
