using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataShare;

// Generated from example definition: specification/datashare/resource-manager/Microsoft.DataShare/stable/2021-08-01/examples/ProviderShareSubscriptions_Adjust.json
// this example is just showing the usage of "ProviderShareSubscriptions_Adjust" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ProviderShareSubscriptionResource created on azure
// for more information of creating ProviderShareSubscriptionResource, please refer to the document of ProviderShareSubscriptionResource
string subscriptionId = "12345678-1234-1234-12345678abc";
string resourceGroupName = "SampleResourceGroup";
string accountName = "Account1";
string shareName = "Share1";
string providerShareSubscriptionId = "4256e2cf-0f82-4865-961b-12f83333f487";
ResourceIdentifier providerShareSubscriptionResourceId = ProviderShareSubscriptionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, shareName, providerShareSubscriptionId);
ProviderShareSubscriptionResource providerShareSubscription = client.GetProviderShareSubscriptionResource(providerShareSubscriptionResourceId);

// invoke the operation
ProviderShareSubscriptionData data = new ProviderShareSubscriptionData
{
    ExpireOn = DateTimeOffset.Parse("2020-12-26T22:33:24.5785265Z"),
};
ProviderShareSubscriptionResource result = await providerShareSubscription.AdjustAsync(data);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ProviderShareSubscriptionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
