using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Automanage;
using Azure.ResourceManager.Automanage.Models;

// Generated from example definition: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/getConfigurationProfileHCRPAssignment.json
// this example is just showing the usage of "ConfigurationProfileHCRPAssignments_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmResource created on azure
// for more information of creating ArmResource, please refer to the document of ArmResource

// get the collection of this AutomanageHcrpConfigurationProfileAssignmentResource
string subscriptionId = "mySubscriptionId";
string resourceGroupName = "myResourceGroupName";
string machineName = "myMachineName";
ResourceIdentifier scopeId = new ResourceIdentifier(string.Format("/subscriptions/{0}/resourceGroups/{1}/providers/Microsoft.HybridCompute/machines/{2}", subscriptionId, resourceGroupName, machineName));
AutomanageHcrpConfigurationProfileAssignmentCollection collection = client.GetAutomanageHcrpConfigurationProfileAssignments(scopeId);

// invoke the operation
string configurationProfileAssignmentName = "default";
NullableResponse<AutomanageHcrpConfigurationProfileAssignmentResource> response = await collection.GetIfExistsAsync(configurationProfileAssignmentName);
AutomanageHcrpConfigurationProfileAssignmentResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    AutomanageConfigurationProfileAssignmentData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
