using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.VoiceServices;
using Azure.ResourceManager.VoiceServices.Models;

// Generated from example definition: specification/voiceservices/resource-manager/Microsoft.VoiceServices/stable/2023-01-31/examples/TestLines_Delete.json
// this example is just showing the usage of "TestLines_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VoiceServicesTestLineResource created on azure
// for more information of creating VoiceServicesTestLineResource, please refer to the document of VoiceServicesTestLineResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "testrg";
string communicationsGatewayName = "myname";
string testLineName = "myline";
ResourceIdentifier voiceServicesTestLineResourceId = VoiceServicesTestLineResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, communicationsGatewayName, testLineName);
VoiceServicesTestLineResource voiceServicesTestLine = client.GetVoiceServicesTestLineResource(voiceServicesTestLineResourceId);

// invoke the operation
await voiceServicesTestLine.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
