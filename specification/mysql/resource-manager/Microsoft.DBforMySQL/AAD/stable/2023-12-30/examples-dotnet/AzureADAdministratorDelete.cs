using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MySql.FlexibleServers.Models;
using Azure.ResourceManager.MySql.FlexibleServers;

// Generated from example definition: specification/mysql/resource-manager/Microsoft.DBforMySQL/AAD/stable/2023-12-30/examples/AzureADAdministratorDelete.json
// this example is just showing the usage of "AzureADAdministrators_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MySqlFlexibleServerAadAdministratorResource created on azure
// for more information of creating MySqlFlexibleServerAadAdministratorResource, please refer to the document of MySqlFlexibleServerAadAdministratorResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "testrg";
string serverName = "mysqltestsvc4";
MySqlFlexibleServerAdministratorName administratorName = MySqlFlexibleServerAdministratorName.ActiveDirectory;
ResourceIdentifier mySqlFlexibleServerAadAdministratorResourceId = MySqlFlexibleServerAadAdministratorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, administratorName);
MySqlFlexibleServerAadAdministratorResource mySqlFlexibleServerAadAdministrator = client.GetMySqlFlexibleServerAadAdministratorResource(mySqlFlexibleServerAadAdministratorResourceId);

// invoke the operation
await mySqlFlexibleServerAadAdministrator.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
