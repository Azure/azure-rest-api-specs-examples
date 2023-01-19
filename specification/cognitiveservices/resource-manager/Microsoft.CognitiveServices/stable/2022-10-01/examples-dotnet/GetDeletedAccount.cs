using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.CognitiveServices;

// Generated from example definition: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-10-01/examples/GetDeletedAccount.json
// this example is just showing the usage of "DeletedAccounts_Get" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this CognitiveServicesDeletedAccountResource created on azure
// for more information of creating CognitiveServicesDeletedAccountResource, please refer to the document of CognitiveServicesDeletedAccountResource
string subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
AzureLocation location = new AzureLocation("westus");
string resourceGroupName = "myResourceGroup";
string accountName = "myAccount";
ResourceIdentifier cognitiveServicesDeletedAccountResourceId = CognitiveServicesDeletedAccountResource.CreateResourceIdentifier(subscriptionId, location, resourceGroupName, accountName);
CognitiveServicesDeletedAccountResource cognitiveServicesDeletedAccount = client.GetCognitiveServicesDeletedAccountResource(cognitiveServicesDeletedAccountResourceId);

// invoke the operation
CognitiveServicesDeletedAccountResource result = await cognitiveServicesDeletedAccount.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CognitiveServicesAccountData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
