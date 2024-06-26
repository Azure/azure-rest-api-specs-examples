using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DeploymentManager;

// Generated from example definition: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/servicetopology_get.json
// this example is just showing the usage of "ServiceTopologies_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceTopologyResource created on azure
// for more information of creating ServiceTopologyResource, please refer to the document of ServiceTopologyResource
string subscriptionId = "caac1590-e859-444f-a9e0-62091c0f5929";
string resourceGroupName = "myResourceGroup";
string serviceTopologyName = "myTopology";
ResourceIdentifier serviceTopologyResourceId = ServiceTopologyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceTopologyName);
ServiceTopologyResource serviceTopologyResource = client.GetServiceTopologyResource(serviceTopologyResourceId);

// invoke the operation
ServiceTopologyResource result = await serviceTopologyResource.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ServiceTopologyResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
