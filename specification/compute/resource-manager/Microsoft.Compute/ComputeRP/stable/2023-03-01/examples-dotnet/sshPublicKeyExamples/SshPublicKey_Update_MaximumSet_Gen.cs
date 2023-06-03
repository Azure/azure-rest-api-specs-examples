using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/sshPublicKeyExamples/SshPublicKey_Update_MaximumSet_Gen.json
// this example is just showing the usage of "SshPublicKeys_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SshPublicKeyResource created on azure
// for more information of creating SshPublicKeyResource, please refer to the document of SshPublicKeyResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "rgcompute";
string sshPublicKeyName = "aaaaaaaaaaaa";
ResourceIdentifier sshPublicKeyResourceId = SshPublicKeyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, sshPublicKeyName);
SshPublicKeyResource sshPublicKey = client.GetSshPublicKeyResource(sshPublicKeyResourceId);

// invoke the operation
SshPublicKeyPatch patch = new SshPublicKeyPatch()
{
    PublicKey = "{ssh-rsa public key}",
    Tags =
    {
    ["key2854"] = "a",
    },
};
SshPublicKeyResource result = await sshPublicKey.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SshPublicKeyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
