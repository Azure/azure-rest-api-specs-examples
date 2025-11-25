using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ComputeLimit;

// Generated from example definition: 2025-08-15/GuestSubscriptions_Delete.json
// this example is just showing the usage of "GuestSubscription_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ComputeLimitGuestSubscriptionResource created on azure
// for more information of creating ComputeLimitGuestSubscriptionResource, please refer to the document of ComputeLimitGuestSubscriptionResource
string subscriptionId = "12345678-1234-1234-1234-123456789012";
AzureLocation location = new AzureLocation("eastus");
string guestSubscriptionId = "11111111-1111-1111-1111-111111111111";
ResourceIdentifier computeLimitGuestSubscriptionResourceId = ComputeLimitGuestSubscriptionResource.CreateResourceIdentifier(subscriptionId, location, guestSubscriptionId);
ComputeLimitGuestSubscriptionResource computeLimitGuestSubscription = client.GetComputeLimitGuestSubscriptionResource(computeLimitGuestSubscriptionResourceId);

// invoke the operation
await computeLimitGuestSubscription.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
