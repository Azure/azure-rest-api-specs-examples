using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Analysis.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Analysis;

// Generated from example definition: specification/analysisservices/resource-manager/Microsoft.AnalysisServices/stable/2017-08-01/examples/updateServer.json
// this example is just showing the usage of "Servers_Update" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this AnalysisServerResource created on azure
// for more information of creating AnalysisServerResource, please refer to the document of AnalysisServerResource
string subscriptionId = "613192d7-503f-477a-9cfe-4efc3ee2bd60";
string resourceGroupName = "TestRG";
string serverName = "azsdktest";
ResourceIdentifier analysisServerResourceId = AnalysisServerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName);
AnalysisServerResource analysisServer = client.GetAnalysisServerResource(analysisServerResourceId);

// invoke the operation
AnalysisServerPatch patch = new AnalysisServerPatch()
{
    Sku = new AnalysisResourceSku("S1")
    {
        Tier = AnalysisSkuTier.Standard,
        Capacity = 1,
    },
    Tags =
    {
    ["testKey"] = "testValue",
    },
    AsAdministratorIdentities =
    {
    "azsdktest@microsoft.com","azsdktest2@microsoft.com"
    },
};
ArmOperation<AnalysisServerResource> lro = await analysisServer.UpdateAsync(WaitUntil.Completed, patch);
AnalysisServerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AnalysisServerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
