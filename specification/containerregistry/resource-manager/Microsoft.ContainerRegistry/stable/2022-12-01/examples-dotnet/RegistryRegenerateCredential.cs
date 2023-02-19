using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ContainerRegistry;
using Azure.ResourceManager.ContainerRegistry.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2022-12-01/examples/RegistryRegenerateCredential.json
// this example is just showing the usage of "Registries_RegenerateCredential" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerRegistryResource created on azure
// for more information of creating ContainerRegistryResource, please refer to the document of ContainerRegistryResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string registryName = "myRegistry";
ResourceIdentifier containerRegistryResourceId = ContainerRegistryResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, registryName);
ContainerRegistryResource containerRegistry = client.GetContainerRegistryResource(containerRegistryResourceId);

// invoke the operation
ContainerRegistryCredentialRegenerateContent content = new ContainerRegistryCredentialRegenerateContent(ContainerRegistryPasswordName.Password);
ContainerRegistryListCredentialsResult result = await containerRegistry.RegenerateCredentialAsync(content);

Console.WriteLine($"Succeeded: {result}");
