using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DevTestLabs;
using Azure.ResourceManager.DevTestLabs.Models;

// Generated from example definition: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Environments_Get.json
// this example is just showing the usage of "Environments_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DevTestLabUserResource created on azure
// for more information of creating DevTestLabUserResource, please refer to the document of DevTestLabUserResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "resourceGroupName";
string labName = "{labName}";
string userName = "@me";
ResourceIdentifier devTestLabUserResourceId = DevTestLabUserResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, labName, userName);
DevTestLabUserResource devTestLabUser = client.GetDevTestLabUserResource(devTestLabUserResourceId);

// get the collection of this DevTestLabEnvironmentResource
DevTestLabEnvironmentCollection collection = devTestLabUser.GetDevTestLabEnvironments();

// invoke the operation
string name = "{environmentName}";
bool result = await collection.ExistsAsync(name);

Console.WriteLine($"Succeeded: {result}");
