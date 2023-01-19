using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DevTestLabs;
using Azure.ResourceManager.DevTestLabs.Models;

// Generated from example definition: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ServiceFabrics_Get.json
// this example is just showing the usage of "ServiceFabrics_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DevTestLabServiceFabricResource created on azure
// for more information of creating DevTestLabServiceFabricResource, please refer to the document of DevTestLabServiceFabricResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "resourceGroupName";
string labName = "{labName}";
string userName = "{userName}";
string name = "{serviceFabricName}";
ResourceIdentifier devTestLabServiceFabricResourceId = DevTestLabServiceFabricResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, labName, userName, name);
DevTestLabServiceFabricResource devTestLabServiceFabric = client.GetDevTestLabServiceFabricResource(devTestLabServiceFabricResourceId);

// invoke the operation
DevTestLabServiceFabricResource result = await devTestLabServiceFabric.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DevTestLabServiceFabricData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
