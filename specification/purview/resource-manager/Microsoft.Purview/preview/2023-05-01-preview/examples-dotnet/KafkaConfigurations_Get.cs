using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Purview;
using Azure.ResourceManager.Purview.Models;

// Generated from example definition: specification/purview/resource-manager/Microsoft.Purview/preview/2023-05-01-preview/examples/KafkaConfigurations_Get.json
// this example is just showing the usage of "KafkaConfigurations_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PurviewAccountResource created on azure
// for more information of creating PurviewAccountResource, please refer to the document of PurviewAccountResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "rgpurview";
string accountName = "account1";
ResourceIdentifier purviewAccountResourceId = PurviewAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
PurviewAccountResource purviewAccount = client.GetPurviewAccountResource(purviewAccountResourceId);

// get the collection of this PurviewKafkaConfigurationResource
PurviewKafkaConfigurationCollection collection = purviewAccount.GetPurviewKafkaConfigurations();

// invoke the operation
string kafkaConfigurationName = "kafkaConfigName";
NullableResponse<PurviewKafkaConfigurationResource> response = await collection.GetIfExistsAsync(kafkaConfigurationName);
PurviewKafkaConfigurationResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    PurviewKafkaConfigurationData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
