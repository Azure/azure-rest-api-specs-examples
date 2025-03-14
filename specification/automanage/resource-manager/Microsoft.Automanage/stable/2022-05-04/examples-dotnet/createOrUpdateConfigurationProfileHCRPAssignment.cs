using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Automanage.Models;
using Azure.ResourceManager.Automanage;

// Generated from example definition: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/createOrUpdateConfigurationProfileHCRPAssignment.json
// this example is just showing the usage of "ConfigurationProfileHCRPAssignments_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// get the collection of this AutomanageHcrpConfigurationProfileAssignmentResource
string subscriptionId = "mySubscriptionId";
string resourceGroupName = "myResourceGroupName";
string machineName = "myMachineName";
string scope = $"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}";
AutomanageHcrpConfigurationProfileAssignmentCollection collection = client.GetAutomanageHcrpConfigurationProfileAssignments(new ResourceIdentifier(scope));

// invoke the operation
string configurationProfileAssignmentName = "default";
AutomanageConfigurationProfileAssignmentData data = new AutomanageConfigurationProfileAssignmentData
{
    Properties = new AutomanageConfigurationProfileAssignmentProperties
    {
        ConfigurationProfile = new ResourceIdentifier("/providers/Microsoft.Automanage/bestPractices/AzureBestPracticesProduction"),
    },
};
ArmOperation<AutomanageHcrpConfigurationProfileAssignmentResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, configurationProfileAssignmentName, data);
AutomanageHcrpConfigurationProfileAssignmentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AutomanageConfigurationProfileAssignmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
