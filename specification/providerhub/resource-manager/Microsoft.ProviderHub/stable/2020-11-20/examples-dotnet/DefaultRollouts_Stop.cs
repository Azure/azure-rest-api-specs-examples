using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ProviderHub.Models;
using Azure.ResourceManager.ProviderHub;

// Generated from example definition: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/DefaultRollouts_Stop.json
// this example is just showing the usage of "DefaultRollouts_Stop" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DefaultRolloutResource created on azure
// for more information of creating DefaultRolloutResource, please refer to the document of DefaultRolloutResource
string subscriptionId = "ab7a8701-f7ef-471a-a2f4-d0ebbf494f77";
string providerNamespace = "Microsoft.Contoso";
string rolloutName = "2020week10";
ResourceIdentifier defaultRolloutResourceId = DefaultRolloutResource.CreateResourceIdentifier(subscriptionId, providerNamespace, rolloutName);
DefaultRolloutResource defaultRollout = client.GetDefaultRolloutResource(defaultRolloutResourceId);

// invoke the operation
await defaultRollout.StopAsync();

Console.WriteLine("Succeeded");
