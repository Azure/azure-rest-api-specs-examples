using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MySql.FlexibleServers;
using Azure.ResourceManager.MySql.FlexibleServers.Models;

// Generated from example definition: specification/mysql/resource-manager/Microsoft.DBforMySQL/AAD/preview/2023-06-01-preview/examples/AzureADAdministratorGet.json
// this example is just showing the usage of "AzureADAdministrators_Get" operation, for the dependent resources, they will have to be created separately.

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
MySqlFlexibleServerAadAdministratorResource result = await mySqlFlexibleServerAadAdministrator.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MySqlFlexibleServerAadAdministratorData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
