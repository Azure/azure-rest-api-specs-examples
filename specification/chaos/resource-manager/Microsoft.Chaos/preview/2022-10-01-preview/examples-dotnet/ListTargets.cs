using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Chaos;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/chaos/resource-manager/Microsoft.Chaos/preview/2022-10-01-preview/examples/ListTargets.json
// this example is just showing the usage of "Targets_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "6b052e15-03d3-4f17-b2e1-be7f07588291";
string resourceGroupName = "exampleRG";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this TargetResource
string parentProviderNamespace = "Microsoft.Compute";
string parentResourceType = "virtualMachines";
string parentResourceName = "exampleVM";
TargetCollection collection = resourceGroupResource.GetTargets(parentProviderNamespace, parentResourceType, parentResourceName);

// invoke the operation and iterate over the result
string continuationToken = null;
await foreach (TargetResource item in collection.GetAllAsync(continuationToken: continuationToken))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    TargetData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
