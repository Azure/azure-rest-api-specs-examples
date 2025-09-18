using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.WorkloadOrchestration.Models;
using Azure.ResourceManager.WorkloadOrchestration;

// Generated from example definition: 2025-06-01/SolutionTemplateVersions_Get_MaximumSet_Gen.json
// this example is just showing the usage of "SolutionTemplateVersion_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EdgeSolutionTemplateVersionResource created on azure
// for more information of creating EdgeSolutionTemplateVersionResource, please refer to the document of EdgeSolutionTemplateVersionResource
string subscriptionId = "9D54FE4C-00AF-4836-8F48-B6A9C4E47192";
string resourceGroupName = "rgconfigurationmanager";
string solutionTemplateName = "testname";
string solutionTemplateVersionName = "1.0.0";
ResourceIdentifier edgeSolutionTemplateVersionResourceId = EdgeSolutionTemplateVersionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, solutionTemplateName, solutionTemplateVersionName);
EdgeSolutionTemplateVersionResource edgeSolutionTemplateVersion = client.GetEdgeSolutionTemplateVersionResource(edgeSolutionTemplateVersionResourceId);

// invoke the operation
EdgeSolutionTemplateVersionResource result = await edgeSolutionTemplateVersion.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
EdgeSolutionTemplateVersionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
