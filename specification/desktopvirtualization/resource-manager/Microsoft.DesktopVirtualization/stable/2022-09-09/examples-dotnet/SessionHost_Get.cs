using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DesktopVirtualization;

// Generated from example definition: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2022-09-09/examples/SessionHost_Get.json
// this example is just showing the usage of "SessionHosts_Get" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this SessionHostResource
SessionHostCollection collection = hostPool.GetSessionHosts();

// invoke the operation
string sessionHostName = "sessionHost1.microsoft.com";
bool result = await collection.ExistsAsync(sessionHostName);

Console.WriteLine($"Succeeded: {result}");
