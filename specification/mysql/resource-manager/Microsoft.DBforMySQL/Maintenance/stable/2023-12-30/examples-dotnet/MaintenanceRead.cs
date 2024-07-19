using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MySql.FlexibleServers.Models;
using Azure.ResourceManager.MySql.FlexibleServers;

// Generated from example definition: specification/mysql/resource-manager/Microsoft.DBforMySQL/Maintenance/stable/2023-12-30/examples/MaintenanceRead.json
// this example is just showing the usage of "Maintenances_Read" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MySqlFlexibleServerMaintenanceResource created on azure
// for more information of creating MySqlFlexibleServerMaintenanceResource, please refer to the document of MySqlFlexibleServerMaintenanceResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "TestGroup";
string serverName = "testserver";
string maintenanceName = "_T9Q-TS8";
ResourceIdentifier mySqlFlexibleServerMaintenanceResourceId = MySqlFlexibleServerMaintenanceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, maintenanceName);
MySqlFlexibleServerMaintenanceResource mySqlFlexibleServerMaintenance = client.GetMySqlFlexibleServerMaintenanceResource(mySqlFlexibleServerMaintenanceResourceId);

// invoke the operation
MySqlFlexibleServerMaintenanceResource result = await mySqlFlexibleServerMaintenance.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MySqlFlexibleServerMaintenanceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
