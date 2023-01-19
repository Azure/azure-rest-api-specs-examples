using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HealthcareApis;
using Azure.ResourceManager.HealthcareApis.Models;

// Generated from example definition: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2022-06-01/examples/legacy/ServiceDeletePrivateEndpointConnection.json
// this example is just showing the usage of "PrivateEndpointConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HealthcareApisServicePrivateEndpointConnectionResource created on azure
// for more information of creating HealthcareApisServicePrivateEndpointConnectionResource, please refer to the document of HealthcareApisServicePrivateEndpointConnectionResource
string subscriptionId = "subid";
string resourceGroupName = "rgname";
string resourceName = "service1";
string privateEndpointConnectionName = "myConnection";
ResourceIdentifier healthcareApisServicePrivateEndpointConnectionResourceId = HealthcareApisServicePrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, privateEndpointConnectionName);
HealthcareApisServicePrivateEndpointConnectionResource healthcareApisServicePrivateEndpointConnection = client.GetHealthcareApisServicePrivateEndpointConnectionResource(healthcareApisServicePrivateEndpointConnectionResourceId);

// invoke the operation
await healthcareApisServicePrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
