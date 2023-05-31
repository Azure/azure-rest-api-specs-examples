using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Cdn;
using Azure.ResourceManager.Cdn.Models;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Secrets_Create.json
// this example is just showing the usage of "FrontDoorSecrets_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ProfileResource created on azure
// for more information of creating ProfileResource, please refer to the document of ProfileResource
string subscriptionId = "subid";
string resourceGroupName = "RG";
string profileName = "profile1";
ResourceIdentifier profileResourceId = ProfileResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, profileName);
ProfileResource profile = client.GetProfileResource(profileResourceId);

// get the collection of this FrontDoorSecretResource
FrontDoorSecretCollection collection = profile.GetFrontDoorSecrets();

// invoke the operation
string secretName = "secret1";
FrontDoorSecretData data = new FrontDoorSecretData()
{
    Properties = new CustomerCertificateProperties(new WritableSubResource()
    {
        Id = new ResourceIdentifier("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.KeyVault/vault/kvName/secrets/certificatename"),
    })
    {
        SecretVersion = "abcdef1234578900abcdef1234567890",
        UseLatestVersion = false,
    },
};
ArmOperation<FrontDoorSecretResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, secretName, data);
FrontDoorSecretResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FrontDoorSecretData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
