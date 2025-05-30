using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerRegistry.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.ContainerRegistry;

// Generated from example definition: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2025-04-01/examples/CredentialSetCreate.json
// this example is just showing the usage of "CredentialSets_Create" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this ContainerRegistryCredentialSetResource
ContainerRegistryCredentialSetCollection collection = containerRegistry.GetContainerRegistryCredentialSets();

// invoke the operation
string credentialSetName = "myCredentialSet";
ContainerRegistryCredentialSetData data = new ContainerRegistryCredentialSetData
{
    Identity = new ManagedServiceIdentity("SystemAssigned"),
    LoginServer = "docker.io",
    AuthCredentials = {new ContainerRegistryAuthCredential
    {
    Name = ContainerRegistryCredentialName.Credential1,
    UsernameSecretIdentifier = "https://myvault.vault.azure.net/secrets/username",
    PasswordSecretIdentifier = "https://myvault.vault.azure.net/secrets/password",
    }},
};
ArmOperation<ContainerRegistryCredentialSetResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, credentialSetName, data);
ContainerRegistryCredentialSetResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerRegistryCredentialSetData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
