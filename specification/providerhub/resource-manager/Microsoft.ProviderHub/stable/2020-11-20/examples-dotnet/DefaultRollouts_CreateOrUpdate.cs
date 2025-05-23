using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ProviderHub.Models;
using Azure.ResourceManager.ProviderHub;

// Generated from example definition: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/DefaultRollouts_CreateOrUpdate.json
// this example is just showing the usage of "DefaultRollouts_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DefaultRolloutResource created on azure
// for more information of creating DefaultRolloutResource, please refer to the document of DefaultRolloutResource
string subscriptionId = "ab7a8701-f7ef-471a-a2f4-d0ebbf494f77";
string providerNamespace = "Microsoft.Contoso";
string rolloutName = "2020week10";
ResourceIdentifier defaultRolloutResourceId = DefaultRolloutResource.CreateResourceIdentifier(subscriptionId, providerNamespace, rolloutName);
DefaultRolloutResource defaultRollout = client.GetDefaultRolloutResource(defaultRolloutResourceId);

// invoke the operation
DefaultRolloutData data = new DefaultRolloutData
{
    Properties = new DefaultRolloutProperties
    {
        Specification = new DefaultRolloutSpecification
        {
            Canary = new CanaryTrafficRegionRolloutConfiguration
            {
                SkipRegions = { new AzureLocation("eastus2euap") },
            },
            RestOfTheWorldGroupTwo = new TrafficRegionRolloutConfiguration
            {
                WaitDuration = XmlConvert.ToTimeSpan("PT4H"),
            },
        },
    },
};
ArmOperation<DefaultRolloutResource> lro = await defaultRollout.UpdateAsync(WaitUntil.Completed, data);
DefaultRolloutResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DefaultRolloutData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
