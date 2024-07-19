using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HealthcareApis.Models;
using Azure.ResourceManager.HealthcareApis;

// Generated from example definition: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/iotconnectors/iotconnector_fhirdestination_Delete.json
// this example is just showing the usage of "IotConnectorFhirDestination_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HealthcareApisIotFhirDestinationResource created on azure
// for more information of creating HealthcareApisIotFhirDestinationResource, please refer to the document of HealthcareApisIotFhirDestinationResource
string subscriptionId = "subid";
string resourceGroupName = "testRG";
string workspaceName = "workspace1";
string iotConnectorName = "blue";
string fhirDestinationName = "dest1";
ResourceIdentifier healthcareApisIotFhirDestinationResourceId = HealthcareApisIotFhirDestinationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, iotConnectorName, fhirDestinationName);
HealthcareApisIotFhirDestinationResource healthcareApisIotFhirDestination = client.GetHealthcareApisIotFhirDestinationResource(healthcareApisIotFhirDestinationResourceId);

// invoke the operation
await healthcareApisIotFhirDestination.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
