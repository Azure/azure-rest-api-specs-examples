using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ServiceFabricManagedClusters.Models;

// Generated from example definition: specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/stable/2022-01-01/examples/managedUnsupportedVMSizesGet_example.json
// this example is just showing the usage of "managedUnsupportedVMSizes_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
AzureLocation location = new AzureLocation("eastus");
string vmSize = "Standard_B1ls1";
ServiceFabricManagedUnsupportedVmSize result = await subscriptionResource.GetManagedUnsupportedVmSizeAsync(location, vmSize);

Console.WriteLine($"Succeeded: {result}");
