using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementGetWorkspaceApiDiagnostic.json
// this example is just showing the usage of "WorkspaceApiDiagnostic_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceWorkspaceApiDiagnosticResource created on azure
// for more information of creating ServiceWorkspaceApiDiagnosticResource, please refer to the document of ServiceWorkspaceApiDiagnosticResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string workspaceId = "wks1";
string apiId = "57d1f7558aa04f15146d9d8a";
string diagnosticId = "applicationinsights";
ResourceIdentifier serviceWorkspaceApiDiagnosticResourceId = ServiceWorkspaceApiDiagnosticResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, workspaceId, apiId, diagnosticId);
ServiceWorkspaceApiDiagnosticResource serviceWorkspaceApiDiagnostic = client.GetServiceWorkspaceApiDiagnosticResource(serviceWorkspaceApiDiagnosticResourceId);

// invoke the operation
ServiceWorkspaceApiDiagnosticResource result = await serviceWorkspaceApiDiagnostic.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DiagnosticContractData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
