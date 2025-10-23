using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MongoCluster.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.MongoCluster;

// Generated from example definition: 2025-09-01/MongoClusters_NameAvailability.json
// this example is just showing the usage of "MongoClusters_CheckMongoClusterNameAvailability" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
AzureLocation location = new AzureLocation("westus2");
MongoClusterNameAvailabilityContent content = new MongoClusterNameAvailabilityContent
{
    Name = "newmongocluster",
    ResourceType = "Microsoft.DocumentDB/mongoClusters",
};
MongoClusterNameAvailabilityResult result = await subscriptionResource.CheckMongoClusterNameAvailabilityAsync(location, content);

Console.WriteLine($"Succeeded: {result}");
