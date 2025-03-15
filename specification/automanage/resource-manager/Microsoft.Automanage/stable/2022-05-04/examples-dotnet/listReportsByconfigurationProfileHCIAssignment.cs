using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Automanage;

// Generated from example definition: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/listReportsByconfigurationProfileHCIAssignment.json
// this example is just showing the usage of "HCIReports_ListByConfigurationProfileAssignments" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this AutomanageHciClusterConfigurationProfileAssignmentReportResource
AutomanageHciClusterConfigurationProfileAssignmentReportCollection collection = automanageHciClusterConfigurationProfileAssignment.GetAutomanageHciClusterConfigurationProfileAssignmentReports();

// invoke the operation and iterate over the result
await foreach (AutomanageHciClusterConfigurationProfileAssignmentReportResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    AutomanageConfigurationProfileAssignmentReportData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
