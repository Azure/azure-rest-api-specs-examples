using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DevTestLabs;
using Azure.ResourceManager.DevTestLabs.Models;

// Generated from example definition: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/CustomImages_Delete.json
// this example is just showing the usage of "CustomImages_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DevTestLabCustomImageResource created on azure
// for more information of creating DevTestLabCustomImageResource, please refer to the document of DevTestLabCustomImageResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "resourceGroupName";
string labName = "{labName}";
string name = "{customImageName}";
ResourceIdentifier devTestLabCustomImageResourceId = DevTestLabCustomImageResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, labName, name);
DevTestLabCustomImageResource devTestLabCustomImage = client.GetDevTestLabCustomImageResource(devTestLabCustomImageResourceId);

// invoke the operation
await devTestLabCustomImage.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
