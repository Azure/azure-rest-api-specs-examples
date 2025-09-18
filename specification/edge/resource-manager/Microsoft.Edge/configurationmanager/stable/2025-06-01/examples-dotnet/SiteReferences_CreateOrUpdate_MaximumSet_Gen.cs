using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.WorkloadOrchestration.Models;
using Azure.ResourceManager.WorkloadOrchestration;

// Generated from example definition: 2025-06-01/SiteReferences_CreateOrUpdate_MaximumSet_Gen.json
// this example is just showing the usage of "SiteReference_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EdgeContextResource created on azure
// for more information of creating EdgeContextResource, please refer to the document of EdgeContextResource
string subscriptionId = "9D54FE4C-00AF-4836-8F48-B6A9C4E47192";
string resourceGroupName = "rgconfigurationmanager";
string contextName = "testname";
ResourceIdentifier edgeContextResourceId = EdgeContextResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, contextName);
EdgeContextResource edgeContext = client.GetEdgeContextResource(edgeContextResourceId);

// get the collection of this EdgeSiteReferenceResource
EdgeSiteReferenceCollection collection = edgeContext.GetEdgeSiteReferences();

// invoke the operation
string siteReferenceName = "testname";
EdgeSiteReferenceData data = new EdgeSiteReferenceData
{
    Properties = new EdgeSiteReferenceProperties("xxjpxdcaumewwgpbwzkcrgrcw"),
};
ArmOperation<EdgeSiteReferenceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, siteReferenceName, data);
EdgeSiteReferenceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
EdgeSiteReferenceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
