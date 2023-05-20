using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.VoiceServices;
using Azure.ResourceManager.VoiceServices.Models;

// Generated from example definition: specification/voiceservices/resource-manager/Microsoft.VoiceServices/stable/2023-01-31/examples/TestLines_CreateOrUpdate.json
// this example is just showing the usage of "TestLines_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VoiceServicesCommunicationsGatewayResource created on azure
// for more information of creating VoiceServicesCommunicationsGatewayResource, please refer to the document of VoiceServicesCommunicationsGatewayResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "testrg";
string communicationsGatewayName = "myname";
ResourceIdentifier voiceServicesCommunicationsGatewayResourceId = VoiceServicesCommunicationsGatewayResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, communicationsGatewayName);
VoiceServicesCommunicationsGatewayResource voiceServicesCommunicationsGateway = client.GetVoiceServicesCommunicationsGatewayResource(voiceServicesCommunicationsGatewayResourceId);

// get the collection of this VoiceServicesTestLineResource
VoiceServicesTestLineCollection collection = voiceServicesCommunicationsGateway.GetVoiceServicesTestLines();

// invoke the operation
string testLineName = "myline";
VoiceServicesTestLineData data = new VoiceServicesTestLineData(new AzureLocation("useast"))
{
    PhoneNumber = "+1-555-1234",
    Purpose = VoiceServicesTestLinePurpose.Automated,
};
ArmOperation<VoiceServicesTestLineResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, testLineName, data);
VoiceServicesTestLineResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VoiceServicesTestLineData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
