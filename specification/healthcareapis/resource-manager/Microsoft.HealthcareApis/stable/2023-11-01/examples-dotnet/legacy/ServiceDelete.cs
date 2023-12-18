using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HealthcareApis;
using Azure.ResourceManager.HealthcareApis.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2023-11-01/examples/legacy/ServiceDelete.json
// this example is just showing the usage of "Services_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HealthcareApisServiceResource created on azure
// for more information of creating HealthcareApisServiceResource, please refer to the document of HealthcareApisServiceResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string resourceName = "service1";
ResourceIdentifier healthcareApisServiceResourceId = HealthcareApisServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
HealthcareApisServiceResource healthcareApisService = client.GetHealthcareApisServiceResource(healthcareApisServiceResourceId);

// invoke the operation
await healthcareApisService.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
