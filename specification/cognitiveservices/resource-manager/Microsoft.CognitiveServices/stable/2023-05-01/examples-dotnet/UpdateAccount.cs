using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CognitiveServices;
using Azure.ResourceManager.CognitiveServices.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2023-05-01/examples/UpdateAccount.json
// this example is just showing the usage of "Accounts_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CognitiveServicesAccountResource created on azure
// for more information of creating CognitiveServicesAccountResource, please refer to the document of CognitiveServicesAccountResource
string subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
string resourceGroupName = "bvttest";
string accountName = "bingSearch";
ResourceIdentifier cognitiveServicesAccountResourceId = CognitiveServicesAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
CognitiveServicesAccountResource cognitiveServicesAccount = client.GetCognitiveServicesAccountResource(cognitiveServicesAccountResourceId);

// invoke the operation
CognitiveServicesAccountData data = new CognitiveServicesAccountData(new AzureLocation("global"))
{
    Sku = new CognitiveServicesSku("S2"),
};
ArmOperation<CognitiveServicesAccountResource> lro = await cognitiveServicesAccount.UpdateAsync(WaitUntil.Completed, data);
CognitiveServicesAccountResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CognitiveServicesAccountData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
