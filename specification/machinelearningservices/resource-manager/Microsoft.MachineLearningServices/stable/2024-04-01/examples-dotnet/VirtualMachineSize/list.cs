using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MachineLearning.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.MachineLearning;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/VirtualMachineSize/list.json
// this example is just showing the usage of "VirtualMachineSizes_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation and iterate over the result
AzureLocation location = new AzureLocation("eastus");
await foreach (MachineLearningVmSize item in subscriptionResource.GetMachineLearningVmSizesAsync(location))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
