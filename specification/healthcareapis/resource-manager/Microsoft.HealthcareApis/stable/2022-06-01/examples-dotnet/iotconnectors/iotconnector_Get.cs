using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HealthcareApis;
using Azure.ResourceManager.HealthcareApis.Models;
using Azure.ResourceManager.Models;

// Generated from example definition: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2022-06-01/examples/iotconnectors/iotconnector_Get.json
// this example is just showing the usage of "IotConnectors_Get" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
HealthcareApisIotConnectorResource result = await healthcareApisIotConnector.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HealthcareApisIotConnectorData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
