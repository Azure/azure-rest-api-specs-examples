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

// this example assumes you already have this AutomanageHcrpConfigurationProfileAssignmentResource created on azure
// for more information of creating AutomanageHcrpConfigurationProfileAssignmentResource, please refer to the document of AutomanageHcrpConfigurationProfileAssignmentResource
string subscriptionId = "mySubscriptionId";
string resourceGroupName = "myResourceGroupName";
string machineName = "myMachineName";
string configurationProfileAssignmentName = "default";
ResourceIdentifier automanageHcrpConfigurationProfileAssignmentResourceId = AutomanageHcrpConfigurationProfileAssignmentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, machineName, configurationProfileAssignmentName);
AutomanageHcrpConfigurationProfileAssignmentResource automanageHcrpConfigurationProfileAssignment = client.GetAutomanageHcrpConfigurationProfileAssignmentResource(automanageHcrpConfigurationProfileAssignmentResourceId);

// invoke the operation
AutomanageHcrpConfigurationProfileAssignmentResource result = await automanageHcrpConfigurationProfileAssignment.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AutomanageConfigurationProfileAssignmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
