using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerRegistry.Models;
using Azure.ResourceManager.ContainerRegistry;

// Generated from example definition: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/Registry/stable/2025-11-01/examples/WebhookGet.json
// this example is just showing the usage of "Webhooks_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerRegistryWebhookResource created on azure
// for more information of creating ContainerRegistryWebhookResource, please refer to the document of ContainerRegistryWebhookResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string registryName = "myRegistry";
string webhookName = "myWebhook";
ResourceIdentifier containerRegistryWebhookResourceId = ContainerRegistryWebhookResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, registryName, webhookName);
ContainerRegistryWebhookResource containerRegistryWebhook = client.GetContainerRegistryWebhookResource(containerRegistryWebhookResourceId);

// invoke the operation
ContainerRegistryWebhookResource result = await containerRegistryWebhook.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerRegistryWebhookData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
