using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CognitiveServices.Models;
using Azure.ResourceManager.CognitiveServices;

// Generated from example definition: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-10-01/examples/PutDeployment.json
// this example is just showing the usage of "Deployments_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this CognitiveServicesAccountResource created on azure
// for more information of creating CognitiveServicesAccountResource, please refer to the document of CognitiveServicesAccountResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "resourceGroupName";
string accountName = "accountName";
ResourceIdentifier cognitiveServicesAccountResourceId = CognitiveServicesAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
CognitiveServicesAccountResource cognitiveServicesAccount = client.GetCognitiveServicesAccountResource(cognitiveServicesAccountResourceId);

// get the collection of this CognitiveServicesAccountDeploymentResource
CognitiveServicesAccountDeploymentCollection collection = cognitiveServicesAccount.GetCognitiveServicesAccountDeployments();

// invoke the operation
string deploymentName = "deploymentName";
CognitiveServicesAccountDeploymentData data = new CognitiveServicesAccountDeploymentData()
{
    Properties = new CognitiveServicesAccountDeploymentProperties()
    {
        Model = new CognitiveServicesAccountDeploymentModel()
        {
            Format = "OpenAI",
            Name = "ada",
            Version = "1",
        },
        ScaleSettings = new CognitiveServicesAccountDeploymentScaleSettings()
        {
            ScaleType = CognitiveServicesAccountDeploymentScaleType.Manual,
            Capacity = 1,
        },
    },
};
ArmOperation<CognitiveServicesAccountDeploymentResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, deploymentName, data);
CognitiveServicesAccountDeploymentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CognitiveServicesAccountDeploymentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
