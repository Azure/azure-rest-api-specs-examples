using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CognitiveServices.Models;
using Azure.ResourceManager.CognitiveServices;

// Generated from example definition: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-06-01/examples/AccountCapabilityHost/createOrUpdate.json
// this example is just showing the usage of "AccountCapabilityHosts_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CognitiveServicesCapabilityHostResource created on azure
// for more information of creating CognitiveServicesCapabilityHostResource, please refer to the document of CognitiveServicesCapabilityHostResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string accountName = "account-1";
string capabilityHostName = "capabilityHostName";
ResourceIdentifier cognitiveServicesCapabilityHostResourceId = CognitiveServicesCapabilityHostResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, capabilityHostName);
CognitiveServicesCapabilityHostResource cognitiveServicesCapabilityHost = client.GetCognitiveServicesCapabilityHostResource(cognitiveServicesCapabilityHostResourceId);

// invoke the operation
CognitiveServicesCapabilityHostData data = new CognitiveServicesCapabilityHostData(new CognitiveServicesCapabilityHostProperties
{
    CustomerSubnet = "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroups/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet",
});
ArmOperation<CognitiveServicesCapabilityHostResource> lro = await cognitiveServicesCapabilityHost.UpdateAsync(WaitUntil.Completed, data);
CognitiveServicesCapabilityHostResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CognitiveServicesCapabilityHostData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
