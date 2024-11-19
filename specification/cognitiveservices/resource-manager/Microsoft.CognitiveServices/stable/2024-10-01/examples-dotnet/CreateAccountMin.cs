using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CognitiveServices.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.CognitiveServices;

// Generated from example definition: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2024-10-01/examples/CreateAccountMin.json
// this example is just showing the usage of "Accounts_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
string resourceGroupName = "myResourceGroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this CognitiveServicesAccountResource
CognitiveServicesAccountCollection collection = resourceGroupResource.GetCognitiveServicesAccounts();

// invoke the operation
string accountName = "testCreate1";
CognitiveServicesAccountData data = new CognitiveServicesAccountData(new AzureLocation("West US"))
{
    Kind = "CognitiveServices",
    Sku = new CognitiveServicesSku("S0"),
    Identity = new ManagedServiceIdentity("SystemAssigned"),
    Properties = new CognitiveServicesAccountProperties(),
};
ArmOperation<CognitiveServicesAccountResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, accountName, data);
CognitiveServicesAccountResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CognitiveServicesAccountData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
