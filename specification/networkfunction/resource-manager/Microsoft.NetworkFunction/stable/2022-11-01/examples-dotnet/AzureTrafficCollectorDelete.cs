using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NetworkFunction;
using Azure.ResourceManager.NetworkFunction.Models;

// Generated from example definition: specification/networkfunction/resource-manager/Microsoft.NetworkFunction/stable/2022-11-01/examples/AzureTrafficCollectorDelete.json
// this example is just showing the usage of "AzureTrafficCollectors_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AzureTrafficCollectorResource created on azure
// for more information of creating AzureTrafficCollectorResource, please refer to the document of AzureTrafficCollectorResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string azureTrafficCollectorName = "atc";
ResourceIdentifier azureTrafficCollectorResourceId = AzureTrafficCollectorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, azureTrafficCollectorName);
AzureTrafficCollectorResource azureTrafficCollector = client.GetAzureTrafficCollectorResource(azureTrafficCollectorResourceId);

// invoke the operation
await azureTrafficCollector.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
