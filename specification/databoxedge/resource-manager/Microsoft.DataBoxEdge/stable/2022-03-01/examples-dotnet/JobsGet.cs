using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataBoxEdge;

// Generated from example definition: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/JobsGet.json
// this example is just showing the usage of "Jobs_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataBoxEdgeJobResource created on azure
// for more information of creating DataBoxEdgeJobResource, please refer to the document of DataBoxEdgeJobResource
string subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
string resourceGroupName = "GroupForEdgeAutomation";
string deviceName = "testedgedevice";
string name = "159a00c7-8543-4343-9435-263ac87df3bb";
ResourceIdentifier dataBoxEdgeJobResourceId = DataBoxEdgeJobResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, deviceName, name);
DataBoxEdgeJobResource dataBoxEdgeJob = client.GetDataBoxEdgeJobResource(dataBoxEdgeJobResourceId);

// invoke the operation
DataBoxEdgeJobResource result = await dataBoxEdgeJob.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DataBoxEdgeJobData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
