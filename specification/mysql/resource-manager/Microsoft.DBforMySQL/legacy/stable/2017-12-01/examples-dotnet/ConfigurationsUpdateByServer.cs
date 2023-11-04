using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MySql;
using Azure.ResourceManager.MySql.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/mysql/resource-manager/Microsoft.DBforMySQL/legacy/stable/2017-12-01/examples/ConfigurationsUpdateByServer.json
// this example is just showing the usage of "ServerParameters_ListUpdateConfigurations" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MySqlServerResource created on azure
// for more information of creating MySqlServerResource, please refer to the document of MySqlServerResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "testrg";
string serverName = "mysqltestsvc1";
ResourceIdentifier mySqlServerResourceId = MySqlServerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName);
MySqlServerResource mySqlServer = client.GetMySqlServerResource(mySqlServerResourceId);

// invoke the operation
MySqlConfigurations value = new MySqlConfigurations();
ArmOperation<MySqlConfigurations> lro = await mySqlServer.UpdateConfigurationsAsync(WaitUntil.Completed, value);
MySqlConfigurations result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
