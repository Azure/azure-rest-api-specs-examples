using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.WorkloadsSapVirtualInstance.Models;
using Azure.ResourceManager.WorkloadsSapVirtualInstance;

// Generated from example definition: 2024-09-01/SapVirtualInstances_InvokeAvailabilityZoneDetails_northeurope.json
// this example is just showing the usage of "SAPVirtualInstances_GetAvailabilityZoneDetails" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "8e17e36c-42e9-4cd5-a078-7b44883414e0";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
AzureLocation location = new AzureLocation("northeurope");
SapAvailabilityZoneDetailsContent content = new SapAvailabilityZoneDetailsContent(new AzureLocation("northeurope"), SapProductType.S4Hana, SapDatabaseType.Hana);
SapAvailabilityZoneDetailsResult result = await subscriptionResource.GetAvailabilityZoneDetailsSapVirtualInstanceAsync(location, content);

Console.WriteLine($"Succeeded: {result}");
