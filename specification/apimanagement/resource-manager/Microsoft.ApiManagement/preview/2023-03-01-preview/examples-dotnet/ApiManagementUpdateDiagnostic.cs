using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-03-01-preview/examples/ApiManagementUpdateDiagnostic.json
// this example is just showing the usage of "Diagnostic_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementDiagnosticResource created on azure
// for more information of creating ApiManagementDiagnosticResource, please refer to the document of ApiManagementDiagnosticResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string diagnosticId = "applicationinsights";
ResourceIdentifier apiManagementDiagnosticResourceId = ApiManagementDiagnosticResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, diagnosticId);
ApiManagementDiagnosticResource apiManagementDiagnostic = client.GetApiManagementDiagnosticResource(apiManagementDiagnosticResourceId);

// invoke the operation
ETag ifMatch = new ETag("*");
DiagnosticContractData data = new DiagnosticContractData
{
    AlwaysLog = AlwaysLog.AllErrors,
    LoggerId = "/loggers/applicationinsights",
    Sampling = new SamplingSettings
    {
        SamplingType = SamplingType.Fixed,
        Percentage = 50,
    },
    Frontend = new PipelineDiagnosticSettings
    {
        Request = new HttpMessageDiagnostic
        {
            Headers = { "Content-type" },
            BodyBytes = 512,
        },
        Response = new HttpMessageDiagnostic
        {
            Headers = { "Content-type" },
            BodyBytes = 512,
        },
    },
    Backend = new PipelineDiagnosticSettings
    {
        Request = new HttpMessageDiagnostic
        {
            Headers = { "Content-type" },
            BodyBytes = 512,
        },
        Response = new HttpMessageDiagnostic
        {
            Headers = { "Content-type" },
            BodyBytes = 512,
        },
    },
};
ApiManagementDiagnosticResource result = await apiManagementDiagnostic.UpdateAsync(ifMatch, data);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DiagnosticContractData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
