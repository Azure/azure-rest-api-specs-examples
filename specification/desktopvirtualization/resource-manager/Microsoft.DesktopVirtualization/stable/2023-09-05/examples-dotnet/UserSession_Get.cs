using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DesktopVirtualization;

// Generated from example definition: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/UserSession_Get.json
// this example is just showing the usage of "UserSessions_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SessionHostResource created on azure
// for more information of creating SessionHostResource, please refer to the document of SessionHostResource
string subscriptionId = "daefabc0-95b4-48b3-b645-8a753a63c4fa";
string resourceGroupName = "resourceGroup1";
string hostPoolName = "hostPool1";
string sessionHostName = "sessionHost1.microsoft.com";
ResourceIdentifier sessionHostResourceId = SessionHostResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, hostPoolName, sessionHostName);
SessionHostResource sessionHost = client.GetSessionHostResource(sessionHostResourceId);

// get the collection of this UserSessionResource
UserSessionCollection collection = sessionHost.GetUserSessions();

// invoke the operation
string userSessionId = "1";
NullableResponse<UserSessionResource> response = await collection.GetIfExistsAsync(userSessionId);
UserSessionResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    UserSessionData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
