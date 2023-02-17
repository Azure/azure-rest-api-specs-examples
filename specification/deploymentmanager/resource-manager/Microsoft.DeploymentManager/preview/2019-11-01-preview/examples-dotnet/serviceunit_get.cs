using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DeploymentManager;
using Azure.ResourceManager.DeploymentManager.Models;

// Generated from example definition: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/serviceunit_get.json
// this example is just showing the usage of "ServiceUnits_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceResource created on azure
// for more information of creating ServiceResource, please refer to the document of ServiceResource
string subscriptionId = "caac1590-e859-444f-a9e0-62091c0f5929";
string resourceGroupName = "myResourceGroup";
string serviceTopologyName = "myTopology";
string serviceName = "myService";
ResourceIdentifier serviceResourceId = ServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceTopologyName, serviceName);
ServiceResource serviceResource = client.GetServiceResource(serviceResourceId);

// get the collection of this ServiceUnitResource
ServiceUnitResourceCollection collection = serviceResource.GetServiceUnitResources();

// invoke the operation
string serviceUnitName = "myServiceUnit";
bool result = await collection.ExistsAsync(serviceUnitName);

Console.WriteLine($"Succeeded: {result}");
