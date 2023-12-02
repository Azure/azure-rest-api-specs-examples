using System;
using System.Net;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.SqlVirtualMachine;
using Azure.ResourceManager.SqlVirtualMachine.Models;

// Generated from example definition: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/stable/2022-02-01/examples/GetAvailabilityGroupListener.json
// this example is just showing the usage of "AvailabilityGroupListeners_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlVmGroupResource created on azure
// for more information of creating SqlVmGroupResource, please refer to the document of SqlVmGroupResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "testrg";
string sqlVmGroupName = "testvmgroup";
ResourceIdentifier sqlVmGroupResourceId = SqlVmGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, sqlVmGroupName);
SqlVmGroupResource sqlVmGroup = client.GetSqlVmGroupResource(sqlVmGroupResourceId);

// get the collection of this AvailabilityGroupListenerResource
AvailabilityGroupListenerCollection collection = sqlVmGroup.GetAvailabilityGroupListeners();

// invoke the operation
string availabilityGroupListenerName = "agl-test";
NullableResponse<AvailabilityGroupListenerResource> response = await collection.GetIfExistsAsync(availabilityGroupListenerName);
AvailabilityGroupListenerResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    AvailabilityGroupListenerData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
