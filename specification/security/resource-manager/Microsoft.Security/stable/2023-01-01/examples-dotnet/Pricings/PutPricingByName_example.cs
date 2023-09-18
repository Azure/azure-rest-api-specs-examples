using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.SecurityCenter;
using Azure.ResourceManager.SecurityCenter.Models;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/stable/2023-01-01/examples/Pricings/PutPricingByName_example.json
// this example is just showing the usage of "Pricings_Update" operation, for the dependent resources, they will have to be created separately.

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
string pricingName = "CloudPosture";
SecurityCenterPricingData data = new SecurityCenterPricingData()
{
    PricingTier = SecurityCenterPricingTier.Standard,
};
ArmOperation<SecurityCenterPricingResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, pricingName, data);
SecurityCenterPricingResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecurityCenterPricingData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
