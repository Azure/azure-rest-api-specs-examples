using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/sshPublicKeyExamples/SshPublicKey_Create.json
// this example is just showing the usage of "SshPublicKeys_Create" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this SshPublicKeyResource
SshPublicKeyCollection collection = resourceGroupResource.GetSshPublicKeys();

// invoke the operation
string sshPublicKeyName = "mySshPublicKeyName";
SshPublicKeyData data = new SshPublicKeyData(new AzureLocation("westus"))
{
    PublicKey = "{ssh-rsa public key}",
};
ArmOperation<SshPublicKeyResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, sshPublicKeyName, data);
SshPublicKeyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SshPublicKeyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
