using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DesktopVirtualization;
using Azure.ResourceManager.DesktopVirtualization.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/UserSession_ListByHostPool.json
// this example is just showing the usage of "UserSessions_ListByHostPool" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HostPoolResource created on azure
// for more information of creating HostPoolResource, please refer to the document of HostPoolResource
string subscriptionId = "daefabc0-95b4-48b3-b645-8a753a63c4fa";
string resourceGroupName = "resourceGroup1";
string hostPoolName = "hostPool1";
ResourceIdentifier hostPoolResourceId = HostPoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, hostPoolName);
HostPoolResource hostPool = client.GetHostPoolResource(hostPoolResourceId);

// invoke the operation and iterate over the result
string filter = "userPrincipalName eq 'user1@microsoft.com' and state eq 'active'";
int? pageSize = 10;
bool? isDescending = true;
int? initialSkip = 0;
await foreach (UserSessionResource item in hostPool.GetUserSessionsAsync(filter: filter, pageSize: pageSize, isDescending: isDescending, initialSkip: initialSkip))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    UserSessionData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
