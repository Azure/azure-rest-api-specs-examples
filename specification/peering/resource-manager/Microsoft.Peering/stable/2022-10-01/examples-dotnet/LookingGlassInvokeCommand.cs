using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Peering.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Peering;

// Generated from example definition: specification/peering/resource-manager/Microsoft.Peering/stable/2022-10-01/examples/LookingGlassInvokeCommand.json
// this example is just showing the usage of "LookingGlass_Invoke" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "subId";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
LookingGlassCommand command = LookingGlassCommand.Traceroute;
LookingGlassSourceType sourceType = LookingGlassSourceType.AzureRegion;
string sourceLocation = "West US";
string destinationIP = "0.0.0.0";
LookingGlassOutput result = await subscriptionResource.InvokeLookingGlassAsync(command, sourceType, sourceLocation, destinationIP);

Console.WriteLine($"Succeeded: {result}");
