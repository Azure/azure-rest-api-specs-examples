using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Cdn.Models;
using Azure.ResourceManager.Cdn;

// Generated from example definition: specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/AFDOriginGroups_Delete.json
// this example is just showing the usage of "FrontDoorOriginGroups_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FrontDoorOriginGroupResource created on azure
// for more information of creating FrontDoorOriginGroupResource, please refer to the document of FrontDoorOriginGroupResource
string subscriptionId = "subid";
string resourceGroupName = "RG";
string profileName = "profile1";
string originGroupName = "origingroup1";
ResourceIdentifier frontDoorOriginGroupResourceId = FrontDoorOriginGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, profileName, originGroupName);
FrontDoorOriginGroupResource frontDoorOriginGroup = client.GetFrontDoorOriginGroupResource(frontDoorOriginGroupResourceId);

// invoke the operation
await frontDoorOriginGroup.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
