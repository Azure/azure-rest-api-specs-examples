using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridNetwork;
using Azure.ResourceManager.HybridNetwork.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/NetworkFunctionsExecuteRequest.json
// this example is just showing the usage of "NetworkFunctions_ExecuteRequest" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkFunctionResource created on azure
// for more information of creating NetworkFunctionResource, please refer to the document of NetworkFunctionResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string networkFunctionName = "testNetworkfunction";
ResourceIdentifier networkFunctionResourceId = NetworkFunctionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkFunctionName);
NetworkFunctionResource networkFunction = client.GetNetworkFunctionResource(networkFunctionResourceId);

// invoke the operation
ExecuteRequestContent content = new ExecuteRequestContent("serviceEndpoint", new RequestMetadata("/simProfiles/testSimProfile", HttpMethod.Post, "{\"subscriptionProfile\":\"ChantestSubscription15\",\"permanentKey\":\"00112233445566778899AABBCCDDEEFF\",\"opcOperatorCode\":\"63bfa50ee6523365ff14c1f45f88737d\",\"staticIpAddresses\":{\"internet\":{\"ipv4Addr\":\"198.51.100.1\",\"ipv6Prefix\":\"2001:db8:abcd:12::0/64\"},\"another_network\":{\"ipv6Prefix\":\"2001:111:cdef:22::0/64\"}}}")
{
    ApiVersion = "apiVersionQueryString",
});
await networkFunction.ExecuteRequestAsync(WaitUntil.Completed, content);

Console.WriteLine($"Succeeded");
