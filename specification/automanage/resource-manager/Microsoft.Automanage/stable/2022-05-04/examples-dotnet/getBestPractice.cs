using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Automanage;

// Generated from example definition: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/getBestPractice.json
// this example is just showing the usage of "BestPractices_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AutomanageBestPracticeResource created on azure
// for more information of creating AutomanageBestPracticeResource, please refer to the document of AutomanageBestPracticeResource
string bestPracticeName = "azureBestPracticesProduction";
ResourceIdentifier automanageBestPracticeResourceId = AutomanageBestPracticeResource.CreateResourceIdentifier(bestPracticeName);
AutomanageBestPracticeResource automanageBestPractice = client.GetAutomanageBestPracticeResource(automanageBestPracticeResourceId);

// invoke the operation
AutomanageBestPracticeResource result = await automanageBestPractice.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AutomanageBestPracticeData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
