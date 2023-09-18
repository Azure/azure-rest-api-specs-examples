using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.SecurityCenter;
using Azure.ResourceManager.SecurityCenter.Models;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/stable/2023-01-01/examples/Pricings/GetPricingByNameStorageAccounts_example.json
// this example is just showing the usage of "Pricings_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// get the collection of this SecurityCenterPricingResource
SecurityCenterPricingCollection collection = subscriptionResource.GetSecurityCenterPricings();

// invoke the operation
string pricingName = "StorageAccounts";
bool result = await collection.ExistsAsync(pricingName);

Console.WriteLine($"Succeeded: {result}");
