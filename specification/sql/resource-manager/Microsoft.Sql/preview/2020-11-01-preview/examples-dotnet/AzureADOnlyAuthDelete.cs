using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;
using Azure.ResourceManager.Sql.Models;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/AzureADOnlyAuthDelete.json
// this example is just showing the usage of "ServerAzureADOnlyAuthentications_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlServerAzureADOnlyAuthenticationResource created on azure
// for more information of creating SqlServerAzureADOnlyAuthenticationResource, please refer to the document of SqlServerAzureADOnlyAuthenticationResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "sqlcrudtest-4799";
string serverName = "sqlcrudtest-6440";
AuthenticationName authenticationName = AuthenticationName.Default;
ResourceIdentifier sqlServerAzureADOnlyAuthenticationResourceId = SqlServerAzureADOnlyAuthenticationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, authenticationName);
SqlServerAzureADOnlyAuthenticationResource sqlServerAzureADOnlyAuthentication = client.GetSqlServerAzureADOnlyAuthenticationResource(sqlServerAzureADOnlyAuthenticationResourceId);

// invoke the operation
await sqlServerAzureADOnlyAuthentication.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
