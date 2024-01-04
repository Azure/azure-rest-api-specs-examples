using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Nginx;

// Generated from example definition: specification/nginx/resource-manager/NGINX.NGINXPLUS/stable/2023-04-01/examples/Configurations_Get.json
// this example is just showing the usage of "Configurations_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NginxDeploymentResource created on azure
// for more information of creating NginxDeploymentResource, please refer to the document of NginxDeploymentResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string deploymentName = "myDeployment";
ResourceIdentifier nginxDeploymentResourceId = NginxDeploymentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, deploymentName);
NginxDeploymentResource nginxDeployment = client.GetNginxDeploymentResource(nginxDeploymentResourceId);

// get the collection of this NginxConfigurationResource
NginxConfigurationCollection collection = nginxDeployment.GetNginxConfigurations();

// invoke the operation
string configurationName = "default";
NullableResponse<NginxConfigurationResource> response = await collection.GetIfExistsAsync(configurationName);
NginxConfigurationResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    NginxConfigurationData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
