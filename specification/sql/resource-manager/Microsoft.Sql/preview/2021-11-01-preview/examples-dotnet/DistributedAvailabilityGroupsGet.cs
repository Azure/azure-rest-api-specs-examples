using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/DistributedAvailabilityGroupsGet.json
// this example is just showing the usage of "DistributedAvailabilityGroups_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedInstanceResource created on azure
// for more information of creating ManagedInstanceResource, please refer to the document of ManagedInstanceResource
string subscriptionId = "f2669dff-5f08-45dd-b857-b2a60b72cdc9";
string resourceGroupName = "testrg";
string managedInstanceName = "testcl";
ResourceIdentifier managedInstanceResourceId = ManagedInstanceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, managedInstanceName);
ManagedInstanceResource managedInstance = client.GetManagedInstanceResource(managedInstanceResourceId);

// get the collection of this DistributedAvailabilityGroupResource
DistributedAvailabilityGroupCollection collection = managedInstance.GetDistributedAvailabilityGroups();

// invoke the operation
string distributedAvailabilityGroupName = "dag";
NullableResponse<DistributedAvailabilityGroupResource> response = await collection.GetIfExistsAsync(distributedAvailabilityGroupName);
DistributedAvailabilityGroupResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    DistributedAvailabilityGroupData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
