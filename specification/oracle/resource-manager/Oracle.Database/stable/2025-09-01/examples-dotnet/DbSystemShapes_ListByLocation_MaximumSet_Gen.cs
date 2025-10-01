using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.OracleDatabase;

// Generated from example definition: 2025-09-01/DbSystemShapes_ListByLocation_MaximumSet_Gen.json
// this example is just showing the usage of "DbSystemShape_ListByLocation" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// get the collection of this OracleDBSystemShapeResource
AzureLocation location = new AzureLocation("eastus");
OracleDBSystemShapeCollection collection = subscriptionResource.GetOracleDBSystemShapes(location);

// invoke the operation and iterate over the result
string zone = "ymedsvqavemtixp";
await foreach (OracleDBSystemShapeResource item in collection.GetAllAsync(zone: zone))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    OracleDBSystemShapeData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
