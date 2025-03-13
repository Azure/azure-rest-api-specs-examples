using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MachineLearning.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.MachineLearning;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Workspace/BatchDeployment/get.json
// this example is just showing the usage of "BatchDeployments_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningBatchEndpointResource created on azure
// for more information of creating MachineLearningBatchEndpointResource, please refer to the document of MachineLearningBatchEndpointResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string workspaceName = "my-aml-workspace";
string endpointName = "testEndpointName";
ResourceIdentifier machineLearningBatchEndpointResourceId = MachineLearningBatchEndpointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, endpointName);
MachineLearningBatchEndpointResource machineLearningBatchEndpoint = client.GetMachineLearningBatchEndpointResource(machineLearningBatchEndpointResourceId);

// get the collection of this MachineLearningBatchDeploymentResource
MachineLearningBatchDeploymentCollection collection = machineLearningBatchEndpoint.GetMachineLearningBatchDeployments();

// invoke the operation
string deploymentName = "testDeploymentName";
NullableResponse<MachineLearningBatchDeploymentResource> response = await collection.GetIfExistsAsync(deploymentName);
MachineLearningBatchDeploymentResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    MachineLearningBatchDeploymentData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
