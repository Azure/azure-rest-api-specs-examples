using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiCenter;

// Generated from example definition: specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/Deployments_List.json
// this example is just showing the usage of "Deployments_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiCenterApiResource created on azure
// for more information of creating ApiCenterApiResource, please refer to the document of ApiCenterApiResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "contoso-resources";
string serviceName = "contoso";
string workspaceName = "default";
string apiName = "echo-api";
ResourceIdentifier apiCenterApiResourceId = ApiCenterApiResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, workspaceName, apiName);
ApiCenterApiResource apiCenterApi = client.GetApiCenterApiResource(apiCenterApiResourceId);

// get the collection of this ApiCenterDeploymentResource
ApiCenterDeploymentCollection collection = apiCenterApi.GetApiCenterDeployments();

// invoke the operation and iterate over the result
await foreach (ApiCenterDeploymentResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ApiCenterDeploymentData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
