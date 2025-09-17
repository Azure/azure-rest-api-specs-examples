using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.WorkloadOrchestration.Models;
using Azure.ResourceManager.WorkloadOrchestration;

// Generated from example definition: 2025-06-01/SiteReferences_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "SiteReference_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EdgeSiteReferenceResource created on azure
// for more information of creating EdgeSiteReferenceResource, please refer to the document of EdgeSiteReferenceResource
string subscriptionId = "9D54FE4C-00AF-4836-8F48-B6A9C4E47192";
string resourceGroupName = "rgconfigurationmanager";
string contextName = "testname";
string siteReferenceName = "testname";
ResourceIdentifier edgeSiteReferenceResourceId = EdgeSiteReferenceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, contextName, siteReferenceName);
EdgeSiteReferenceResource edgeSiteReference = client.GetEdgeSiteReferenceResource(edgeSiteReferenceResourceId);

// invoke the operation
await edgeSiteReference.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
