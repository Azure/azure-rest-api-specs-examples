using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ServiceGroups.Models;
using Azure.ResourceManager.ServiceGroups;

// Generated from example definition: specification/management/resource-manager/Microsoft.Management/ServiceGroups/preview/2024-02-01-preview/examples/ServiceGroup_Put.json
// this example is just showing the usage of "ServiceGroups_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// get the collection of this ServiceGroupResource
ServiceGroupCollection collection = tenantResource.GetServiceGroups();

// invoke the operation
string serviceGroupName = "ServiceGroup1";
ServiceGroupData data = new ServiceGroupData
{
    Properties = new ServiceGroupProperties
    {
        DisplayName = "ServiceGroup 1 Name",
        ParentResourceId = new ResourceIdentifier("/providers/Microsoft.Management/serviceGroups/RootGroup"),
    },
};
ArmOperation<ServiceGroupResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, serviceGroupName, data);
ServiceGroupResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ServiceGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
