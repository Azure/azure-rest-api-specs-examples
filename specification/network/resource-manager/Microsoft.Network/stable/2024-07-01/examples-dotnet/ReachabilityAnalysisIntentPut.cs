using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/ReachabilityAnalysisIntentPut.json
// this example is just showing the usage of "ReachabilityAnalysisIntents_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ReachabilityAnalysisIntentResource created on azure
// for more information of creating ReachabilityAnalysisIntentResource, please refer to the document of ReachabilityAnalysisIntentResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string networkManagerName = "testNetworkManager";
string workspaceName = "testWorkspace";
string reachabilityAnalysisIntentName = "testAnalysisIntentName";
ResourceIdentifier reachabilityAnalysisIntentResourceId = ReachabilityAnalysisIntentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkManagerName, workspaceName, reachabilityAnalysisIntentName);
ReachabilityAnalysisIntentResource reachabilityAnalysisIntent = client.GetReachabilityAnalysisIntentResource(reachabilityAnalysisIntentResourceId);

// invoke the operation
ReachabilityAnalysisIntentData data = new ReachabilityAnalysisIntentData(new ReachabilityAnalysisIntentProperties(new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/testVmSrc"), new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/testVmDest"), new NetworkVerifierIPTraffic(new string[] { "10.4.0.0" }, new string[] { "10.4.0.1" }, new string[] { "0" }, new string[] { "0" }, new NetworkProtocol[] { NetworkProtocol.Any }))
{
    Description = "A sample reachability analysis intent",
});
ArmOperation<ReachabilityAnalysisIntentResource> lro = await reachabilityAnalysisIntent.UpdateAsync(WaitUntil.Completed, data);
ReachabilityAnalysisIntentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ReachabilityAnalysisIntentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
