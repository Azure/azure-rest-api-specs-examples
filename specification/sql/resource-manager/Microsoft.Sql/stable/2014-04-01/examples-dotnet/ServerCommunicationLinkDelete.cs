using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/ServerCommunicationLinkDelete.json
// this example is just showing the usage of "ServerCommunicationLinks_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlServerCommunicationLinkResource created on azure
// for more information of creating SqlServerCommunicationLinkResource, please refer to the document of SqlServerCommunicationLinkResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "sqlcrudtest-7398";
string serverName = "sqlcrudtest-4645";
string communicationLinkName = "link1";
ResourceIdentifier sqlServerCommunicationLinkResourceId = SqlServerCommunicationLinkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, communicationLinkName);
SqlServerCommunicationLinkResource sqlServerCommunicationLink = client.GetSqlServerCommunicationLinkResource(sqlServerCommunicationLinkResourceId);

// invoke the operation
await sqlServerCommunicationLink.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
