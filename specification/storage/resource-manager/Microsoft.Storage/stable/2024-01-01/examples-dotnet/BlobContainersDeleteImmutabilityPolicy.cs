using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Storage;

// Generated from example definition: specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/BlobContainersDeleteImmutabilityPolicy.json
// this example is just showing the usage of "BlobContainers_DeleteImmutabilityPolicy" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ImmutabilityPolicyResource created on azure
// for more information of creating ImmutabilityPolicyResource, please refer to the document of ImmutabilityPolicyResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "res1581";
string accountName = "sto9621";
string containerName = "container4910";
ResourceIdentifier immutabilityPolicyResourceId = ImmutabilityPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, containerName);
ImmutabilityPolicyResource immutabilityPolicy = client.GetImmutabilityPolicyResource(immutabilityPolicyResourceId);

// invoke the operation
ETag ifMatch = new ETag("8d59f81a7fa7be0");
ArmOperation<ImmutabilityPolicyResource> lro = await immutabilityPolicy.DeleteAsync(WaitUntil.Completed, ifMatch);
ImmutabilityPolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ImmutabilityPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
