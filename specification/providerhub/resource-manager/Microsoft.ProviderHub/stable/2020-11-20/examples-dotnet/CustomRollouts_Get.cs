using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ProviderHub;
using Azure.ResourceManager.ProviderHub.Models;

// Generated from example definition: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/CustomRollouts_Get.json
// this example is just showing the usage of "CustomRollouts_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CustomRolloutResource created on azure
// for more information of creating CustomRolloutResource, please refer to the document of CustomRolloutResource
string subscriptionId = "ab7a8701-f7ef-471a-a2f4-d0ebbf494f77";
string providerNamespace = "Microsoft.Contoso";
string rolloutName = "canaryTesting99";
ResourceIdentifier customRolloutResourceId = CustomRolloutResource.CreateResourceIdentifier(subscriptionId, providerNamespace, rolloutName);
CustomRolloutResource customRollout = client.GetCustomRolloutResource(customRolloutResourceId);

// invoke the operation
CustomRolloutResource result = await customRollout.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CustomRolloutData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
