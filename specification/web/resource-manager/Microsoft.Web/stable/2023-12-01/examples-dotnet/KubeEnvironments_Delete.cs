using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/KubeEnvironments_Delete.json
// this example is just showing the usage of "KubeEnvironments_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this KubeEnvironmentResource created on azure
// for more information of creating KubeEnvironmentResource, please refer to the document of KubeEnvironmentResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "examplerg";
string name = "examplekenv";
ResourceIdentifier kubeEnvironmentResourceId = KubeEnvironmentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name);
KubeEnvironmentResource kubeEnvironment = client.GetKubeEnvironmentResource(kubeEnvironmentResourceId);

// invoke the operation
await kubeEnvironment.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
