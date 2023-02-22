using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Automanage;
using Azure.ResourceManager.Automanage.Models;

// Generated from example definition: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/deleteConfigurationProfileHCIAssignment.json
// this example is just showing the usage of "ConfigurationProfileHCIAssignments_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AutomanageHciClusterConfigurationProfileAssignmentResource created on azure
// for more information of creating AutomanageHciClusterConfigurationProfileAssignmentResource, please refer to the document of AutomanageHciClusterConfigurationProfileAssignmentResource
string subscriptionId = "mySubscriptionId";
string resourceGroupName = "myResourceGroupName";
string clusterName = "myClusterName";
string configurationProfileAssignmentName = "default";
ResourceIdentifier automanageHciClusterConfigurationProfileAssignmentResourceId = AutomanageHciClusterConfigurationProfileAssignmentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, configurationProfileAssignmentName);
AutomanageHciClusterConfigurationProfileAssignmentResource automanageHciClusterConfigurationProfileAssignment = client.GetAutomanageHciClusterConfigurationProfileAssignmentResource(automanageHciClusterConfigurationProfileAssignmentResourceId);

// invoke the operation
await automanageHciClusterConfigurationProfileAssignment.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
