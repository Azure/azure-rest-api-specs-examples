using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/ManagedInstanceRemoveMaintenanceConfiguration.json
// this example is just showing the usage of "ManagedInstances_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedInstanceResource created on azure
// for more information of creating ManagedInstanceResource, please refer to the document of ManagedInstanceResource
string subscriptionId = "20D7082A-0FC7-4468-82BD-542694D5042B";
string resourceGroupName = "testrg";
string managedInstanceName = "testinstance";
ResourceIdentifier managedInstanceResourceId = ManagedInstanceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, managedInstanceName);
ManagedInstanceResource managedInstance = client.GetManagedInstanceResource(managedInstanceResourceId);

// invoke the operation
ManagedInstancePatch patch = new ManagedInstancePatch
{
    MaintenanceConfigurationId = new ResourceIdentifier("/subscriptions/20d7082a-0fc7-4468-82bd-542694d5042b/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_Default"),
};
ArmOperation<ManagedInstanceResource> lro = await managedInstance.UpdateAsync(WaitUntil.Completed, patch);
ManagedInstanceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagedInstanceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
