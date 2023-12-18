using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HealthcareApis;
using Azure.ResourceManager.HealthcareApis.Models;

// Generated from example definition: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2023-11-01/examples/dicomservices/DicomServices_Delete.json
// this example is just showing the usage of "DicomServices_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DicomServiceResource created on azure
// for more information of creating DicomServiceResource, please refer to the document of DicomServiceResource
string subscriptionId = "subid";
string resourceGroupName = "testRG";
string workspaceName = "workspace1";
string dicomServiceName = "blue";
ResourceIdentifier dicomServiceResourceId = DicomServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, dicomServiceName);
DicomServiceResource dicomService = client.GetDicomServiceResource(dicomServiceResourceId);

// invoke the operation
await dicomService.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
