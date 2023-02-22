using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Automanage;
using Azure.ResourceManager.Automanage.Models;

// Generated from example definition: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/deleteConfigurationProfileAssignment.json
// this example is just showing the usage of "ConfigurationProfileAssignments_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AutomanageVmConfigurationProfileAssignmentResource created on azure
// for more information of creating AutomanageVmConfigurationProfileAssignmentResource, please refer to the document of AutomanageVmConfigurationProfileAssignmentResource
string subscriptionId = "mySubscriptionId";
string resourceGroupName = "myResourceGroupName";
string vmName = "myVMName";
string configurationProfileAssignmentName = "default";
ResourceIdentifier automanageVmConfigurationProfileAssignmentResourceId = AutomanageVmConfigurationProfileAssignmentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vmName, configurationProfileAssignmentName);
AutomanageVmConfigurationProfileAssignmentResource automanageVmConfigurationProfileAssignment = client.GetAutomanageVmConfigurationProfileAssignmentResource(automanageVmConfigurationProfileAssignmentResourceId);

// invoke the operation
await automanageVmConfigurationProfileAssignment.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
