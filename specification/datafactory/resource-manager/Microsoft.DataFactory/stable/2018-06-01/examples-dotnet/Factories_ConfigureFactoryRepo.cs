using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Core.Expressions.DataFactory;
using Azure.Identity;
using Azure.ResourceManager.DataFactory.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.DataFactory;

// Generated from example definition: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Factories_ConfigureFactoryRepo.json
// this example is just showing the usage of "Factories_ConfigureFactoryRepo" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "12345678-1234-1234-1234-12345678abc";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
AzureLocation locationId = new AzureLocation("East US");
FactoryRepoContent content = new FactoryRepoContent()
{
    FactoryResourceId = new ResourceIdentifier("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName"),
    RepoConfiguration = new FactoryVstsConfiguration("ADF", "repo", "master", "/", "project")
    {
        TenantId = Guid.Parse(""),
        LastCommitId = "",
    },
};
DataFactoryResource result = await subscriptionResource.ConfigureFactoryRepoInformationAsync(locationId, content);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DataFactoryData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
