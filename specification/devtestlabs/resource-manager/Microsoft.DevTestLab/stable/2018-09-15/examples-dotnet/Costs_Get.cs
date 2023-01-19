using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DevTestLabs;
using Azure.ResourceManager.DevTestLabs.Models;

// Generated from example definition: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Costs_Get.json
// this example is just showing the usage of "Costs_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DevTestLabCostResource created on azure
// for more information of creating DevTestLabCostResource, please refer to the document of DevTestLabCostResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "resourceGroupName";
string labName = "{labName}";
string name = "targetCost";
ResourceIdentifier devTestLabCostResourceId = DevTestLabCostResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, labName, name);
DevTestLabCostResource devTestLabCost = client.GetDevTestLabCostResource(devTestLabCostResourceId);

// invoke the operation
DevTestLabCostResource result = await devTestLabCost.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DevTestLabCostData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
