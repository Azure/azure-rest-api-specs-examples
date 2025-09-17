using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.WorkloadOrchestration.Models;
using Azure.ResourceManager.WorkloadOrchestration;

// Generated from example definition: 2025-06-01/SolutionTemplates_RemoveVersion_MaximumSet_Gen.json
// this example is just showing the usage of "SolutionTemplates_RemoveVersion" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EdgeSolutionTemplateResource created on azure
// for more information of creating EdgeSolutionTemplateResource, please refer to the document of EdgeSolutionTemplateResource
string subscriptionId = "9D54FE4C-00AF-4836-8F48-B6A9C4E47192";
string resourceGroupName = "rgconfigurationmanager";
string solutionTemplateName = "testname";
ResourceIdentifier edgeSolutionTemplateResourceId = EdgeSolutionTemplateResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, solutionTemplateName);
EdgeSolutionTemplateResource edgeSolutionTemplate = client.GetEdgeSolutionTemplateResource(edgeSolutionTemplateResourceId);

// invoke the operation
EdgeVersionContent content = new EdgeVersionContent("ghtvdzgmzncaifrnuumg");
await edgeSolutionTemplate.RemoveVersionAsync(WaitUntil.Completed, content);

Console.WriteLine("Succeeded");
