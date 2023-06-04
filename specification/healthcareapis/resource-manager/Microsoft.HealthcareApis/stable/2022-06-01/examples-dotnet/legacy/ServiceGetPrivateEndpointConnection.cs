using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HealthcareApis;
using Azure.ResourceManager.HealthcareApis.Models;

// Generated from example definition: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2022-06-01/examples/legacy/ServiceGetPrivateEndpointConnection.json
// this example is just showing the usage of "PrivateEndpointConnections_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HealthcareApisServiceResource created on azure
// for more information of creating HealthcareApisServiceResource, please refer to the document of HealthcareApisServiceResource
string subscriptionId = "subid";
string resourceGroupName = "rgname";
string resourceName = "service1";
ResourceIdentifier healthcareApisServiceResourceId = HealthcareApisServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
HealthcareApisServiceResource healthcareApisService = client.GetHealthcareApisServiceResource(healthcareApisServiceResourceId);

// get the collection of this HealthcareApisServicePrivateEndpointConnectionResource
HealthcareApisServicePrivateEndpointConnectionCollection collection = healthcareApisService.GetHealthcareApisServicePrivateEndpointConnections();

// invoke the operation
string privateEndpointConnectionName = "myConnection";
bool result = await collection.ExistsAsync(privateEndpointConnectionName);

Console.WriteLine($"Succeeded: {result}");
