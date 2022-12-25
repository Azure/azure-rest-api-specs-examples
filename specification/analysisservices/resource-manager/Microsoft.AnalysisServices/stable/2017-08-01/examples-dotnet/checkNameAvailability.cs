using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Analysis.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Analysis;

// Generated from example definition: specification/analysisservices/resource-manager/Microsoft.AnalysisServices/stable/2017-08-01/examples/checkNameAvailability.json
// this example is just showing the usage of "Servers_CheckNameAvailability" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "613192d7-503f-477a-9cfe-4efc3ee2bd60";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
AzureLocation location = new AzureLocation("West US");
AnalysisServerNameAvailabilityContent content = new AnalysisServerNameAvailabilityContent()
{
    Name = "azsdktest",
    ResourceType = "Microsoft.AnalysisServices/servers",
};
AnalysisServerNameAvailabilityResult result = await subscriptionResource.CheckAnalysisServerNameAvailabilityAsync(location, content);

Console.WriteLine($"Succeeded: {result}");
