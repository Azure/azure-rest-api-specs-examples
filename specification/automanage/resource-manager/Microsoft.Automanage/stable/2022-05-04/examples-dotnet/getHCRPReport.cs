using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Automanage;

// Generated from example definition: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/getHCRPReport.json
// this example is just showing the usage of "HCRPReports_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AutomanageHcrpConfigurationProfileAssignmentResource created on azure
// for more information of creating AutomanageHcrpConfigurationProfileAssignmentResource, please refer to the document of AutomanageHcrpConfigurationProfileAssignmentResource
string subscriptionId = "mySubscriptionId";
string resourceGroupName = "myResourceGroupName";
string machineName = "myMachineName";
string configurationProfileAssignmentName = "default";
ResourceIdentifier automanageHcrpConfigurationProfileAssignmentResourceId = AutomanageHcrpConfigurationProfileAssignmentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, machineName, configurationProfileAssignmentName);
AutomanageHcrpConfigurationProfileAssignmentResource automanageHcrpConfigurationProfileAssignment = client.GetAutomanageHcrpConfigurationProfileAssignmentResource(automanageHcrpConfigurationProfileAssignmentResourceId);

// get the collection of this AutomanageHcrpConfigurationProfileAssignmentReportResource
AutomanageHcrpConfigurationProfileAssignmentReportCollection collection = automanageHcrpConfigurationProfileAssignment.GetAutomanageHcrpConfigurationProfileAssignmentReports();

// invoke the operation
string reportName = "b4e9ee6b-1717-4ff0-a8d2-e6d72c33d5f4";
NullableResponse<AutomanageHcrpConfigurationProfileAssignmentReportResource> response = await collection.GetIfExistsAsync(reportName);
AutomanageHcrpConfigurationProfileAssignmentReportResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    AutomanageConfigurationProfileAssignmentReportData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
