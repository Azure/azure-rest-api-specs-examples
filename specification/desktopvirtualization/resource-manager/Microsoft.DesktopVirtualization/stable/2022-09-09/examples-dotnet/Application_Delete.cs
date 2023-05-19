using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DesktopVirtualization;
using Azure.ResourceManager.DesktopVirtualization.Models;

// Generated from example definition: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2022-09-09/examples/Application_Delete.json
// this example is just showing the usage of "Applications_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualApplicationResource created on azure
// for more information of creating VirtualApplicationResource, please refer to the document of VirtualApplicationResource
string subscriptionId = "daefabc0-95b4-48b3-b645-8a753a63c4fa";
string resourceGroupName = "resourceGroup1";
string applicationGroupName = "applicationGroup1";
string applicationName = "application1";
ResourceIdentifier virtualApplicationResourceId = VirtualApplicationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, applicationGroupName, applicationName);
VirtualApplicationResource virtualApplication = client.GetVirtualApplicationResource(virtualApplicationResourceId);

// invoke the operation
await virtualApplication.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
