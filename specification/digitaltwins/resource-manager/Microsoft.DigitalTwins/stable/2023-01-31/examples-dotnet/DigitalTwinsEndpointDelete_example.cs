using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DigitalTwins;
using Azure.ResourceManager.DigitalTwins.Models;

// Generated from example definition: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2023-01-31/examples/DigitalTwinsEndpointDelete_example.json
// this example is just showing the usage of "DigitalTwinsEndpoint_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DigitalTwinsEndpointResource created on azure
// for more information of creating DigitalTwinsEndpointResource, please refer to the document of DigitalTwinsEndpointResource
string subscriptionId = "50016170-c839-41ba-a724-51e9df440b9e";
string resourceGroupName = "resRg";
string resourceName = "myDigitalTwinsService";
string endpointName = "myendpoint";
ResourceIdentifier digitalTwinsEndpointResourceId = DigitalTwinsEndpointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, endpointName);
DigitalTwinsEndpointResource digitalTwinsEndpointResource = client.GetDigitalTwinsEndpointResource(digitalTwinsEndpointResourceId);

// invoke the operation
ArmOperation<DigitalTwinsEndpointResource> lro = await digitalTwinsEndpointResource.DeleteAsync(WaitUntil.Completed);
DigitalTwinsEndpointResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DigitalTwinsEndpointResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
