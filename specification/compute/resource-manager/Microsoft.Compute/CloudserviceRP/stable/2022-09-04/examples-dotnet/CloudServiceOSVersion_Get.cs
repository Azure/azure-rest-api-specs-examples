using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-09-04/examples/CloudServiceOSVersion_Get.json
// this example is just showing the usage of "CloudServiceOperatingSystems_GetOSVersion" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "{subscription-id}";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// get the collection of this CloudServiceOSVersionResource
AzureLocation location = new AzureLocation("westus2");
CloudServiceOSVersionCollection collection = subscriptionResource.GetCloudServiceOSVersions(location);

// invoke the operation
string osVersionName = "WA-GUEST-OS-3.90_202010-02";
bool result = await collection.ExistsAsync(osVersionName);

Console.WriteLine($"Succeeded: {result}");
