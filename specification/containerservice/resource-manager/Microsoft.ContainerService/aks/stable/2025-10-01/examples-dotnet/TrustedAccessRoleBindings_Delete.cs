using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerService;

// Generated from example definition: specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/TrustedAccessRoleBindings_Delete.json
// this example is just showing the usage of "TrustedAccessRoleBindings_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerServiceTrustedAccessRoleBindingResource created on azure
// for more information of creating ContainerServiceTrustedAccessRoleBindingResource, please refer to the document of ContainerServiceTrustedAccessRoleBindingResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string resourceName = "clustername1";
string trustedAccessRoleBindingName = "binding1";
ResourceIdentifier containerServiceTrustedAccessRoleBindingResourceId = ContainerServiceTrustedAccessRoleBindingResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, trustedAccessRoleBindingName);
ContainerServiceTrustedAccessRoleBindingResource containerServiceTrustedAccessRoleBinding = client.GetContainerServiceTrustedAccessRoleBindingResource(containerServiceTrustedAccessRoleBindingResourceId);

// invoke the operation
await containerServiceTrustedAccessRoleBinding.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
