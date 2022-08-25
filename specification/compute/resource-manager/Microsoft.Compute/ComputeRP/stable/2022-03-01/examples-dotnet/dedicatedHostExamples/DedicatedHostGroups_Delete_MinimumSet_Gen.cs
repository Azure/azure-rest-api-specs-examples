using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/dedicatedHostExamples/DedicatedHostGroups_Delete_MinimumSet_Gen.json
// this example is just showing the usage of "DedicatedHostGroups_Delete" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this DedicatedHostGroupResource created on azure
// for more information of creating DedicatedHostGroupResource, please refer to the document of DedicatedHostGroupResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "rgcompute";
string hostGroupName = "aaaa";
ResourceIdentifier dedicatedHostGroupResourceId = DedicatedHostGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, hostGroupName);
DedicatedHostGroupResource dedicatedHostGroup = client.GetDedicatedHostGroupResource(dedicatedHostGroupResourceId);

// invoke the operation
await dedicatedHostGroup.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
