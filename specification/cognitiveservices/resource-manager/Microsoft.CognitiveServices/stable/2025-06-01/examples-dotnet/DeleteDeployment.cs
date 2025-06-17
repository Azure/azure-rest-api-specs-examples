using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CognitiveServices.Models;
using Azure.ResourceManager.CognitiveServices;

// Generated from example definition: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-06-01/examples/DeleteDeployment.json
// this example is just showing the usage of "Deployments_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CognitiveServicesAccountDeploymentResource created on azure
// for more information of creating CognitiveServicesAccountDeploymentResource, please refer to the document of CognitiveServicesAccountDeploymentResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "resourceGroupName";
string accountName = "accountName";
string deploymentName = "deploymentName";
ResourceIdentifier cognitiveServicesAccountDeploymentResourceId = CognitiveServicesAccountDeploymentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, deploymentName);
CognitiveServicesAccountDeploymentResource cognitiveServicesAccountDeployment = client.GetCognitiveServicesAccountDeploymentResource(cognitiveServicesAccountDeploymentResourceId);

// invoke the operation
await cognitiveServicesAccountDeployment.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
