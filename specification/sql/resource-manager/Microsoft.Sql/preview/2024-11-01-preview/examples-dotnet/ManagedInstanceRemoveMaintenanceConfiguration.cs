using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2024-11-01-preview/examples/ManagedInstanceRemoveMaintenanceConfiguration.json
// this example is just showing the usage of "ManagedInstances_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedInstanceResource created on azure
// for more information of creating ManagedInstanceResource, please refer to the document of ManagedInstanceResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "testrg";
string managedInstanceName = "testinstance";
ResourceIdentifier managedInstanceResourceId = ManagedInstanceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, managedInstanceName);
ManagedInstanceResource managedInstance = client.GetManagedInstanceResource(managedInstanceResourceId);

// invoke the operation
ManagedInstancePatch patch = new ManagedInstancePatch
{
    MaintenanceConfigurationId = new ResourceIdentifier("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_Default"),
};
ArmOperation<ManagedInstanceResource> lro = await managedInstance.UpdateAsync(WaitUntil.Completed, patch);
ManagedInstanceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagedInstanceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
