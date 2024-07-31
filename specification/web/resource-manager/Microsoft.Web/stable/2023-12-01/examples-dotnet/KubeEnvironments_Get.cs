using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/KubeEnvironments_Get.json
// this example is just showing the usage of "KubeEnvironments_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this KubeEnvironmentResource created on azure
// for more information of creating KubeEnvironmentResource, please refer to the document of KubeEnvironmentResource
string subscriptionId = "8efdecc5-919e-44eb-b179-915dca89ebf9";
string resourceGroupName = "examplerg";
string name = "jlaw-demo1";
ResourceIdentifier kubeEnvironmentResourceId = KubeEnvironmentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name);
KubeEnvironmentResource kubeEnvironment = client.GetKubeEnvironmentResource(kubeEnvironmentResourceId);

// invoke the operation
KubeEnvironmentResource result = await kubeEnvironment.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
KubeEnvironmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
