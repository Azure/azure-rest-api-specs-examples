using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppService;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/web/resource-manager/Microsoft.DomainRegistration/stable/2021-02-01/examples/GetTopLevelDomain.json
// this example is just showing the usage of "TopLevelDomains_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// get the collection of this TopLevelDomainResource
TopLevelDomainCollection collection = subscriptionResource.GetTopLevelDomains();

// invoke the operation
string name = "com";
bool result = await collection.ExistsAsync(name);

Console.WriteLine($"Succeeded: {result}");
