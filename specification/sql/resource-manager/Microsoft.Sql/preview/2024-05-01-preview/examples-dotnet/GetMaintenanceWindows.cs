using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2024-05-01-preview/examples/GetMaintenanceWindows.json
// this example is just showing the usage of "MaintenanceWindows_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MaintenanceWindowsResource created on azure
// for more information of creating MaintenanceWindowsResource, please refer to the document of MaintenanceWindowsResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "Default-SQL-SouthEastAsia";
string serverName = "testsvr";
string databaseName = "testdb";
ResourceIdentifier maintenanceWindowsResourceId = MaintenanceWindowsResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, databaseName);
MaintenanceWindowsResource maintenanceWindows = client.GetMaintenanceWindowsResource(maintenanceWindowsResourceId);

// invoke the operation
string maintenanceWindowName = "current";
MaintenanceWindowsResource result = await maintenanceWindows.GetAsync(maintenanceWindowName);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MaintenanceWindowsData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
