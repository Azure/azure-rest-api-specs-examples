using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Nginx.Models;
using Azure.ResourceManager.Nginx;

// Generated from example definition: specification/nginx/resource-manager/NGINX.NGINXPLUS/preview/2024-11-01-preview/examples/Deployments_UpdateSubnet.json
// this example is just showing the usage of "Deployments_Update" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
NginxDeploymentPatch patch = new NginxDeploymentPatch();
ArmOperation<NginxDeploymentResource> lro = await nginxDeployment.UpdateAsync(WaitUntil.Completed, patch);
NginxDeploymentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NginxDeploymentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
