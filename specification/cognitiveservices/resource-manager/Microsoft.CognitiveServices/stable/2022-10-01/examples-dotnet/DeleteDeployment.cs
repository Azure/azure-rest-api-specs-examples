using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CognitiveServices.Models;
using Azure.ResourceManager.CognitiveServices;

// Generated from example definition: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-10-01/examples/DeleteDeployment.json
// this example is just showing the usage of "Deployments_Delete" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this CognitiveServicesAccountDeploymentResource created on azure
// for more information of creating CognitiveServicesAccountDeploymentResource, please refer to the document of CognitiveServicesAccountDeploymentResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "resourceGroupName";
string accountName = "accountName";
string deploymentName = "deploymentName";
ResourceIdentifier cognitiveServicesAccountDeploymentResourceId = CognitiveServicesAccountDeploymentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, deploymentName);
CognitiveServicesAccountDeploymentResource cognitiveServicesAccountDeployment = client.GetCognitiveServicesAccountDeploymentResource(cognitiveServicesAccountDeploymentResourceId);

// invoke the operation
await cognitiveServicesAccountDeployment.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
