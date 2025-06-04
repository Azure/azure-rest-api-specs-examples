using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CognitiveServices.Models;
using Azure.ResourceManager.CognitiveServices;

// Generated from example definition: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/preview/2025-04-01-preview/examples/ProjectConnection/create.json
// this example is just showing the usage of "ProjectConnection_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CognitiveServicesProjectResource created on azure
// for more information of creating CognitiveServicesProjectResource, please refer to the document of CognitiveServicesProjectResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "resourceGroup-1";
string accountName = "account-1";
string projectName = "project-1";
ResourceIdentifier cognitiveServicesProjectResourceId = CognitiveServicesProjectResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, projectName);
CognitiveServicesProjectResource cognitiveServicesProject = client.GetCognitiveServicesProjectResource(cognitiveServicesProjectResourceId);

// get the collection of this CognitiveServicesProjectConnectionResource
CognitiveServicesProjectConnectionCollection collection = cognitiveServicesProject.GetCognitiveServicesProjectConnections();

// invoke the operation
string connectionName = "connection-1";
CognitiveServicesConnectionData data = new CognitiveServicesConnectionData(new NoneAuthTypeConnectionProperties
{
    Category = CognitiveServicesConnectionCategory.ContainerRegistry,
    ExpiryOn = DateTimeOffset.Parse("2024-03-15T14:30:00Z"),
    Target = "[tartget url]",
});
ArmOperation<CognitiveServicesProjectConnectionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, connectionName, data);
CognitiveServicesProjectConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CognitiveServicesConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
