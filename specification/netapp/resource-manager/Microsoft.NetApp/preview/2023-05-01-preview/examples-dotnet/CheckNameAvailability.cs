using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NetApp;
using Azure.ResourceManager.NetApp.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/preview/2023-05-01-preview/examples/CheckNameAvailability.json
// this example is just showing the usage of "NetAppResource_CheckNameAvailability" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
AzureLocation location = new AzureLocation("eastus");
NetAppNameAvailabilityContent content = new NetAppNameAvailabilityContent("accName", new NetAppNameAvailabilityResourceType("netAppAccount"), "myRG");
NetAppCheckAvailabilityResult result = await subscriptionResource.CheckNetAppNameAvailabilityAsync(location, content);

Console.WriteLine($"Succeeded: {result}");
