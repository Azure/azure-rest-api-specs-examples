using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Compute;
// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/dedicatedHostExamples/DedicatedHostGroup_Get_UltraSSDEnabledDedicatedHostGroup.json
// this example is just showing the usage of "DedicatedHostGroups_Get" operation, for the dependent resources, they will have to be created separately.
            
// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());
            
// this example assumes you already have this DedicatedHostGroupResource created on azure
// for more information of creating DedicatedHostGroupResource, please refer to the document of DedicatedHostGroupResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "myResourceGroup";
string hostGroupName = "myDedicatedHostGroup";
ResourceIdentifier dedicatedHostGroupResourceId = DedicatedHostGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, hostGroupName);
DedicatedHostGroupResource dedicatedHostGroup = client.GetDedicatedHostGroupResource(dedicatedHostGroupResourceId);
            
// invoke the operation
DedicatedHostGroupResource result = await dedicatedHostGroup.GetAsync();
            
// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DedicatedHostGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
