using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/BastionShareableLinkDeleteByToken.json
// this example is just showing the usage of "DeleteBastionShareableLinkByToken" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BastionHostResource created on azure
// for more information of creating BastionHostResource, please refer to the document of BastionHostResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string bastionHostName = "bastionhosttenant";
ResourceIdentifier bastionHostResourceId = BastionHostResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, bastionHostName);
BastionHostResource bastionHost = client.GetBastionHostResource(bastionHostResourceId);

// invoke the operation
BastionShareableLinkTokenListContent content = new BastionShareableLinkTokenListContent
{
    Tokens = { "abcd1234-efgh-hijk-5678-abcdefgh1234", "dcba4321-hgfe-kjih-8765-hgfedcba4321" },
};
await bastionHost.DeleteBastionShareableLinkByTokenAsync(WaitUntil.Completed, content);

Console.WriteLine("Succeeded");
