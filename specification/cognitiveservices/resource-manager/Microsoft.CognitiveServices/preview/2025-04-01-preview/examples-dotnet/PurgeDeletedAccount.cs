using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CognitiveServices;

// Generated from example definition: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/preview/2025-04-01-preview/examples/PurgeDeletedAccount.json
// this example is just showing the usage of "DeletedAccounts_Purge" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CognitiveServicesDeletedAccountResource created on azure
// for more information of creating CognitiveServicesDeletedAccountResource, please refer to the document of CognitiveServicesDeletedAccountResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
AzureLocation location = new AzureLocation("westus");
string resourceGroupName = "myResourceGroup";
string accountName = "PropTest01";
ResourceIdentifier cognitiveServicesDeletedAccountResourceId = CognitiveServicesDeletedAccountResource.CreateResourceIdentifier(subscriptionId, location, resourceGroupName, accountName);
CognitiveServicesDeletedAccountResource cognitiveServicesDeletedAccount = client.GetCognitiveServicesDeletedAccountResource(cognitiveServicesDeletedAccountResourceId);

// invoke the operation
await cognitiveServicesDeletedAccount.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
