using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DigitalTwins;
using Azure.ResourceManager.DigitalTwins.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2023-01-31/examples/DigitalTwinsGet_WithIdentity_example.json
// this example is just showing the usage of "DigitalTwins_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "50016170-c839-41ba-a724-51e9df440b9e";
string resourceGroupName = "resRg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this DigitalTwinsDescriptionResource
DigitalTwinsDescriptionCollection collection = resourceGroupResource.GetDigitalTwinsDescriptions();

// invoke the operation
string resourceName = "myDigitalTwinsService";
bool result = await collection.ExistsAsync(resourceName);

Console.WriteLine($"Succeeded: {result}");
