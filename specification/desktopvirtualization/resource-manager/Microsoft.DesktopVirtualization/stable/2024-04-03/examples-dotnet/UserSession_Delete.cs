using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DesktopVirtualization.Models;
using Azure.ResourceManager.DesktopVirtualization;

// Generated from example definition: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2024-04-03/examples/UserSession_Delete.json
// this example is just showing the usage of "UserSessions_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this UserSessionResource created on azure
// for more information of creating UserSessionResource, please refer to the document of UserSessionResource
string subscriptionId = "daefabc0-95b4-48b3-b645-8a753a63c4fa";
string resourceGroupName = "resourceGroup1";
string hostPoolName = "hostPool1";
string sessionHostName = "sessionHost1.microsoft.com";
string userSessionId = "1";
ResourceIdentifier userSessionResourceId = UserSessionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, hostPoolName, sessionHostName, userSessionId);
UserSessionResource userSession = client.GetUserSessionResource(userSessionResourceId);

// invoke the operation
bool? force = true;
await userSession.DeleteAsync(WaitUntil.Completed, force: force);

Console.WriteLine("Succeeded");
