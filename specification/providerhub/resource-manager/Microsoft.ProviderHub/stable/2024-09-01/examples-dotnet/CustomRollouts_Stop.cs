using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ProviderHub.Models;
using Azure.ResourceManager.ProviderHub;

// Generated from example definition: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/CustomRollouts_Stop.json
// this example is just showing the usage of "CustomRollouts_Stop" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CustomRolloutResource created on azure
// for more information of creating CustomRolloutResource, please refer to the document of CustomRolloutResource
string subscriptionId = "ab7a8701-f7ef-471a-a2f4-d0ebbf494f77";
string providerNamespace = "Microsoft.Contoso";
string rolloutName = "2020week10";
ResourceIdentifier customRolloutResourceId = CustomRolloutResource.CreateResourceIdentifier(subscriptionId, providerNamespace, rolloutName);
CustomRolloutResource customRollout = client.GetCustomRolloutResource(customRolloutResourceId);

// invoke the operation
await customRollout.StopAsync();

Console.WriteLine("Succeeded");
