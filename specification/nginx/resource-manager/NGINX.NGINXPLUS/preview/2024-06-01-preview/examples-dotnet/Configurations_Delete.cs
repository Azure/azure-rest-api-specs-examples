using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Nginx.Models;
using Azure.ResourceManager.Nginx;

// Generated from example definition: specification/nginx/resource-manager/NGINX.NGINXPLUS/preview/2024-06-01-preview/examples/Configurations_Delete.json
// this example is just showing the usage of "Configurations_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NginxConfigurationResource created on azure
// for more information of creating NginxConfigurationResource, please refer to the document of NginxConfigurationResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string deploymentName = "myDeployment";
string configurationName = "default";
ResourceIdentifier nginxConfigurationResourceId = NginxConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, deploymentName, configurationName);
NginxConfigurationResource nginxConfiguration = client.GetNginxConfigurationResource(nginxConfigurationResourceId);

// invoke the operation
await nginxConfiguration.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
