using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ResourceHealth;

// Generated from example definition: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2023-10-01-preview/examples/Metadata_GetEntity.json
// this example is just showing the usage of "Metadata_GetEntity" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceHealthMetadataEntityResource created on azure
// for more information of creating ResourceHealthMetadataEntityResource, please refer to the document of ResourceHealthMetadataEntityResource
string name = "status";
ResourceIdentifier resourceHealthMetadataEntityResourceId = ResourceHealthMetadataEntityResource.CreateResourceIdentifier(name);
ResourceHealthMetadataEntityResource resourceHealthMetadataEntity = client.GetResourceHealthMetadataEntityResource(resourceHealthMetadataEntityResourceId);

// invoke the operation
ResourceHealthMetadataEntityResource result = await resourceHealthMetadataEntity.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ResourceHealthMetadataEntityData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
