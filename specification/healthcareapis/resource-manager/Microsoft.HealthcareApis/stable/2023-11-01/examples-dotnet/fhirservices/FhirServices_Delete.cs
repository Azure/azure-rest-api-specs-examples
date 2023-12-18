using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HealthcareApis;
using Azure.ResourceManager.HealthcareApis.Models;

// Generated from example definition: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2023-11-01/examples/fhirservices/FhirServices_Delete.json
// this example is just showing the usage of "FhirServices_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FhirServiceResource created on azure
// for more information of creating FhirServiceResource, please refer to the document of FhirServiceResource
string subscriptionId = "subid";
string resourceGroupName = "testRG";
string workspaceName = "workspace1";
string fhirServiceName = "fhirservice1";
ResourceIdentifier fhirServiceResourceId = FhirServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, fhirServiceName);
FhirServiceResource fhirService = client.GetFhirServiceResource(fhirServiceResourceId);

// invoke the operation
await fhirService.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
