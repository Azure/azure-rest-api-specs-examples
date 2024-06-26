using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/sshPublicKeyExamples/SshPublicKeys_GenerateKeyPair.json
// this example is just showing the usage of "SshPublicKeys_GenerateKeyPair" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this SshPublicKeyResource created on azure
// for more information of creating SshPublicKeyResource, please refer to the document of SshPublicKeyResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string sshPublicKeyName = "mySshPublicKeyName";
ResourceIdentifier sshPublicKeyResourceId = SshPublicKeyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, sshPublicKeyName);
SshPublicKeyResource sshPublicKey = client.GetSshPublicKeyResource(sshPublicKeyResourceId);

// invoke the operation
SshPublicKeyGenerateKeyPairResult result = await sshPublicKey.GenerateKeyPairAsync();

Console.WriteLine($"Succeeded: {result}");
