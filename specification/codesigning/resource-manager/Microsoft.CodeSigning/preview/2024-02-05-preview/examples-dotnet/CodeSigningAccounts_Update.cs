using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.TrustedSigning.Models;
using Azure.ResourceManager.TrustedSigning;

// Generated from example definition: specification/codesigning/resource-manager/Microsoft.CodeSigning/preview/2024-02-05-preview/examples/CodeSigningAccounts_Update.json
// this example is just showing the usage of "CodeSigningAccounts_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TrustedSigningAccountResource created on azure
// for more information of creating TrustedSigningAccountResource, please refer to the document of TrustedSigningAccountResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "MyResourceGroup";
string accountName = "MyAccount";
ResourceIdentifier trustedSigningAccountResourceId = TrustedSigningAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
TrustedSigningAccountResource trustedSigningAccount = client.GetTrustedSigningAccountResource(trustedSigningAccountResourceId);

// invoke the operation
TrustedSigningAccountPatch patch = new TrustedSigningAccountPatch
{
    Tags =
    {
    ["key1"] = "value1"
    },
};
ArmOperation<TrustedSigningAccountResource> lro = await trustedSigningAccount.UpdateAsync(WaitUntil.Completed, patch);
TrustedSigningAccountResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
TrustedSigningAccountData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
