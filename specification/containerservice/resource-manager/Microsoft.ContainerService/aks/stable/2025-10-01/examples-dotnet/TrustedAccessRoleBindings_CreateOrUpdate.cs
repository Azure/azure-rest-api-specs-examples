using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerService;

// Generated from example definition: specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/TrustedAccessRoleBindings_CreateOrUpdate.json
// this example is just showing the usage of "TrustedAccessRoleBindings_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerServiceManagedClusterResource created on azure
// for more information of creating ContainerServiceManagedClusterResource, please refer to the document of ContainerServiceManagedClusterResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string resourceName = "clustername1";
ResourceIdentifier containerServiceManagedClusterResourceId = ContainerServiceManagedClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
ContainerServiceManagedClusterResource containerServiceManagedCluster = client.GetContainerServiceManagedClusterResource(containerServiceManagedClusterResourceId);

// get the collection of this ContainerServiceTrustedAccessRoleBindingResource
ContainerServiceTrustedAccessRoleBindingCollection collection = containerServiceManagedCluster.GetContainerServiceTrustedAccessRoleBindings();

// invoke the operation
string trustedAccessRoleBindingName = "binding1";
ContainerServiceTrustedAccessRoleBindingData data = new ContainerServiceTrustedAccessRoleBindingData(new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/b/providers/Microsoft.MachineLearningServices/workspaces/c"), new string[] { "Microsoft.MachineLearningServices/workspaces/reader", "Microsoft.MachineLearningServices/workspaces/writer" });
ArmOperation<ContainerServiceTrustedAccessRoleBindingResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, trustedAccessRoleBindingName, data);
ContainerServiceTrustedAccessRoleBindingResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerServiceTrustedAccessRoleBindingData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
