using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Storage;

// Generated from example definition: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/BlobContainersExtendImmutabilityPolicy.json
// this example is just showing the usage of "BlobContainers_ExtendImmutabilityPolicy" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ImmutabilityPolicyResource created on azure
// for more information of creating ImmutabilityPolicyResource, please refer to the document of ImmutabilityPolicyResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "res6238";
string accountName = "sto232";
string containerName = "container5023";
ResourceIdentifier immutabilityPolicyResourceId = ImmutabilityPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, containerName);
ImmutabilityPolicyResource immutabilityPolicy = client.GetImmutabilityPolicyResource(immutabilityPolicyResourceId);

// invoke the operation
ETag ifMatch = new ETag("8d59f830d0c3bf9");
ImmutabilityPolicyData data = new ImmutabilityPolicyData()
{
    ImmutabilityPeriodSinceCreationInDays = 100,
};
ImmutabilityPolicyResource result = await immutabilityPolicy.ExtendImmutabilityPolicyAsync(ifMatch, data: data);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ImmutabilityPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
