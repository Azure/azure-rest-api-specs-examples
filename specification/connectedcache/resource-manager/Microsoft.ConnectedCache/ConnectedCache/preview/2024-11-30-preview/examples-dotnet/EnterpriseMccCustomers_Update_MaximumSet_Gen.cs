using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ConnectedCache.Models;
using Azure.ResourceManager.ConnectedCache;

// Generated from example definition: 2024-11-30-preview/EnterpriseMccCustomers_Update_MaximumSet_Gen.json
// this example is just showing the usage of "EnterpriseMccCustomerResource_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EnterpriseMccCustomerResource created on azure
// for more information of creating EnterpriseMccCustomerResource, please refer to the document of EnterpriseMccCustomerResource
string subscriptionId = "12345678-1234-1234-1234-123456789098";
string resourceGroupName = "rgConnectedCache";
string customerResourceName = "MccRPTest1";
ResourceIdentifier enterpriseMccCustomerResourceId = EnterpriseMccCustomerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, customerResourceName);
EnterpriseMccCustomerResource enterpriseMccCustomer = client.GetEnterpriseMccCustomerResource(enterpriseMccCustomerResourceId);

// invoke the operation
ConnectedCachePatchContent content = new ConnectedCachePatchContent
{
    Tags =
    {
    ["key1878"] = "warz"
    },
};
EnterpriseMccCustomerResource result = await enterpriseMccCustomer.UpdateAsync(content);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
EnterpriseMccCustomerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
