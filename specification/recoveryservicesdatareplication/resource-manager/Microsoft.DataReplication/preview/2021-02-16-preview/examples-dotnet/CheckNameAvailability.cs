using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RecoveryServicesDataReplication.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.RecoveryServicesDataReplication;

// Generated from example definition: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/CheckNameAvailability.json
// this example is just showing the usage of "CheckNameAvailability" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "930CEC23-4430-4513-B855-DBA237E2F3BF";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
AzureLocation location = new AzureLocation("trfqtbtmusswpibw");
DataReplicationNameAvailabilityContent content = new DataReplicationNameAvailabilityContent
{
    Name = "updkdcixs",
    ResourceType = new ResourceType("gngmcancdauwhdixjjvqnfkvqc"),
};
DataReplicationNameAvailabilityResult result = await subscriptionResource.CheckDataReplicationNameAvailabilityAsync(location, content: content);

Console.WriteLine($"Succeeded: {result}");
