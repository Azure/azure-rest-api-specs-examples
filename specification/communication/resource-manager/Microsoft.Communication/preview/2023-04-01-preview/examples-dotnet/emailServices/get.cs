using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Communication;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/communication/resource-manager/Microsoft.Communication/preview/2023-04-01-preview/examples/emailServices/get.json
// this example is just showing the usage of "EmailServices_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "11112222-3333-4444-5555-666677778888";
string resourceGroupName = "MyResourceGroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this EmailServiceResource
EmailServiceResourceCollection collection = resourceGroupResource.GetEmailServiceResources();

// invoke the operation
string emailServiceName = "MyEmailServiceResource";
bool result = await collection.ExistsAsync(emailServiceName);

Console.WriteLine($"Succeeded: {result}");
