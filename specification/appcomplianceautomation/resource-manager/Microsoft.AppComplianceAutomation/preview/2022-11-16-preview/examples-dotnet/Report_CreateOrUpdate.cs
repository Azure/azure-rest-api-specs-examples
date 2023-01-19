using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppComplianceAutomation;
using Azure.ResourceManager.AppComplianceAutomation.Models;

// Generated from example definition: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/preview/2022-11-16-preview/examples/Report_CreateOrUpdate.json
// this example is just showing the usage of "Report_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TenantResource created on azure
// for more information of creating TenantResource, please refer to the document of TenantResource
var tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// get the collection of this ReportResource
ReportResourceCollection collection = tenantResource.GetReportResources();

// invoke the operation
string reportName = "testReportName";
ReportResourceData data = new ReportResourceData(new ReportProperties("GMT Standard Time", DateTimeOffset.Parse("2022-03-04T05:11:56.197Z"), new ResourceMetadata[]
{
new ResourceMetadata("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.Network/privateEndpoints/myPrivateEndpoint")
{
Tags =
{
["key1"] = "value1",
},
}
})
{
    OfferGuid = "0000",
});
ArmOperation<ReportResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, reportName, data);
ReportResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ReportResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
