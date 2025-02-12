using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerRegistry.Models;
using Azure.ResourceManager.ContainerRegistry;

// Generated from example definition: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2024-11-01-preview/examples/CredentialSetDelete.json
// this example is just showing the usage of "CredentialSets_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerRegistryCredentialSetResource created on azure
// for more information of creating ContainerRegistryCredentialSetResource, please refer to the document of ContainerRegistryCredentialSetResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string registryName = "myRegistry";
string credentialSetName = "myCredentialSet";
ResourceIdentifier containerRegistryCredentialSetResourceId = ContainerRegistryCredentialSetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, registryName, credentialSetName);
ContainerRegistryCredentialSetResource containerRegistryCredentialSet = client.GetContainerRegistryCredentialSetResource(containerRegistryCredentialSetResourceId);

// invoke the operation
await containerRegistryCredentialSet.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
