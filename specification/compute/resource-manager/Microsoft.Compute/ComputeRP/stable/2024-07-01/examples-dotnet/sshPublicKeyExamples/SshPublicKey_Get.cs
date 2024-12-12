using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/sshPublicKeyExamples/SshPublicKey_Get.json
// this example is just showing the usage of "SshPublicKeys_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SshPublicKeyResource created on azure
// for more information of creating SshPublicKeyResource, please refer to the document of SshPublicKeyResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "myResourceGroup";
string sshPublicKeyName = "mySshPublicKeyName";
ResourceIdentifier sshPublicKeyResourceId = SshPublicKeyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, sshPublicKeyName);
SshPublicKeyResource sshPublicKey = client.GetSshPublicKeyResource(sshPublicKeyResourceId);

// invoke the operation
SshPublicKeyResource result = await sshPublicKey.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SshPublicKeyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
