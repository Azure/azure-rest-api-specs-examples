using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HealthcareApis;
using Azure.ResourceManager.HealthcareApis.Models;

// Generated from example definition: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2023-11-01/examples/iotconnectors/iotconnector_fhirdestination_Get.json
// this example is just showing the usage of "IotConnectorFhirDestination_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HealthcareApisIotConnectorResource created on azure
// for more information of creating HealthcareApisIotConnectorResource, please refer to the document of HealthcareApisIotConnectorResource
string subscriptionId = "subid";
string resourceGroupName = "testRG";
string workspaceName = "workspace1";
string iotConnectorName = "blue";
ResourceIdentifier healthcareApisIotConnectorResourceId = HealthcareApisIotConnectorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, iotConnectorName);
HealthcareApisIotConnectorResource healthcareApisIotConnector = client.GetHealthcareApisIotConnectorResource(healthcareApisIotConnectorResourceId);

// get the collection of this HealthcareApisIotFhirDestinationResource
HealthcareApisIotFhirDestinationCollection collection = healthcareApisIotConnector.GetHealthcareApisIotFhirDestinations();

// invoke the operation
string fhirDestinationName = "dest1";
NullableResponse<HealthcareApisIotFhirDestinationResource> response = await collection.GetIfExistsAsync(fhirDestinationName);
HealthcareApisIotFhirDestinationResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    HealthcareApisIotFhirDestinationData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
