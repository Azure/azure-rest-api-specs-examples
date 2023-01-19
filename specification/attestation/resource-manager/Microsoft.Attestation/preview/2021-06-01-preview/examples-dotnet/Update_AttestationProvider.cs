using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Attestation;
using Azure.ResourceManager.Attestation.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/attestation/resource-manager/Microsoft.Attestation/preview/2021-06-01-preview/examples/Update_AttestationProvider.json
// this example is just showing the usage of "AttestationProviders_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AttestationProviderResource created on azure
// for more information of creating AttestationProviderResource, please refer to the document of AttestationProviderResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "MyResourceGroup";
string providerName = "myattestationprovider";
ResourceIdentifier attestationProviderResourceId = AttestationProviderResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, providerName);
AttestationProviderResource attestationProvider = client.GetAttestationProviderResource(attestationProviderResourceId);

// invoke the operation
AttestationProviderPatch patch = new AttestationProviderPatch()
{
    Tags =
    {
    ["Property1"] = "Value1",
    ["Property2"] = "Value2",
    ["Property3"] = "Value3",
    },
    AttestationServicePatchSpecificParamsPublicNetworkAccess = PublicNetworkAccessType.Disabled,
};
AttestationProviderResource result = await attestationProvider.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AttestationProviderData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
