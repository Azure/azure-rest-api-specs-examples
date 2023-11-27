using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DigitalTwins;

// Generated from example definition: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2023-01-31/examples/PrivateLinkResourcesList_example.json
// this example is just showing the usage of "PrivateLinkResources_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DigitalTwinsDescriptionResource created on azure
// for more information of creating DigitalTwinsDescriptionResource, please refer to the document of DigitalTwinsDescriptionResource
string subscriptionId = "50016170-c839-41ba-a724-51e9df440b9e";
string resourceGroupName = "resRg";
string resourceName = "myDigitalTwinsService";
ResourceIdentifier digitalTwinsDescriptionResourceId = DigitalTwinsDescriptionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
DigitalTwinsDescriptionResource digitalTwinsDescription = client.GetDigitalTwinsDescriptionResource(digitalTwinsDescriptionResourceId);

// get the collection of this DigitalTwinsPrivateLinkResource
DigitalTwinsPrivateLinkResourceCollection collection = digitalTwinsDescription.GetDigitalTwinsPrivateLinkResources();

// invoke the operation and iterate over the result
await foreach (DigitalTwinsPrivateLinkResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    DigitalTwinsPrivateLinkResourceData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
