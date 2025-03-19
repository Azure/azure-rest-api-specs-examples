using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DigitalTwins.Models;
using Azure.ResourceManager.DigitalTwins;

// Generated from example definition: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2023-01-31/examples/DigitalTwinsEndpointPut_WithIdentity_example.json
// this example is just showing the usage of "DigitalTwinsEndpoint_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this DigitalTwinsEndpointResource
DigitalTwinsEndpointResourceCollection collection = digitalTwinsDescription.GetDigitalTwinsEndpointResources();

// invoke the operation
string endpointName = "myServiceBus";
DigitalTwinsEndpointResourceData data = new DigitalTwinsEndpointResourceData(new DigitalTwinsServiceBusProperties
{
    EndpointUri = new Uri("sb://mysb.servicebus.windows.net/"),
    EntityPath = "mysbtopic",
    AuthenticationType = DigitalTwinsAuthenticationType.IdentityBased,
});
ArmOperation<DigitalTwinsEndpointResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, endpointName, data);
DigitalTwinsEndpointResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DigitalTwinsEndpointResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
