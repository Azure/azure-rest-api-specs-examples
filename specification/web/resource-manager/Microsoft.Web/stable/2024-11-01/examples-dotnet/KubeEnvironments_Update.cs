using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService.Models;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/KubeEnvironments_Update.json
// this example is just showing the usage of "KubeEnvironments_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this KubeEnvironmentResource created on azure
// for more information of creating KubeEnvironmentResource, please refer to the document of KubeEnvironmentResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "examplerg";
string name = "testkubeenv";
ResourceIdentifier kubeEnvironmentResourceId = KubeEnvironmentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name);
KubeEnvironmentResource kubeEnvironment = client.GetKubeEnvironmentResource(kubeEnvironmentResourceId);

// invoke the operation
KubeEnvironmentPatch patch = new KubeEnvironmentPatch
{
    StaticIP = "1.2.3.4",
};
KubeEnvironmentResource result = await kubeEnvironment.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
KubeEnvironmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
