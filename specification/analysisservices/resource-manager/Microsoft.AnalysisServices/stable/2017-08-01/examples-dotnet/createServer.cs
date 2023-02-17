using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Analysis;
using Azure.ResourceManager.Analysis.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/analysisservices/resource-manager/Microsoft.AnalysisServices/stable/2017-08-01/examples/createServer.json
// this example is just showing the usage of "Servers_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "613192d7-503f-477a-9cfe-4efc3ee2bd60";
string resourceGroupName = "TestRG";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this AnalysisServerResource
AnalysisServerCollection collection = resourceGroupResource.GetAnalysisServers();

// invoke the operation
string serverName = "azsdktest";
AnalysisServerData data = new AnalysisServerData(new AzureLocation("West US"), new AnalysisResourceSku("S1")
{
    Tier = AnalysisSkuTier.Standard,
    Capacity = 1,
})
{
    AsAdministratorIdentities =
    {
    "azsdktest@microsoft.com","azsdktest2@microsoft.com"
    },
    AnalysisServerSku = new AnalysisResourceSku("S1")
    {
        Tier = AnalysisSkuTier.Standard,
        Capacity = 1,
    },
    Tags =
    {
    ["testKey"] = "testValue",
    },
};
ArmOperation<AnalysisServerResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, serverName, data);
AnalysisServerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AnalysisServerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
