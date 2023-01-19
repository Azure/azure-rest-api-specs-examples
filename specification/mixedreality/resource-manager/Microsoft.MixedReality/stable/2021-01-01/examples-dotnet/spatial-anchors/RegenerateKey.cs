using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MixedReality;
using Azure.ResourceManager.MixedReality.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/mixedreality/resource-manager/Microsoft.MixedReality/stable/2021-01-01/examples/spatial-anchors/RegenerateKey.json
// this example is just showing the usage of "SpatialAnchorsAccounts_RegenerateKeys" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SpatialAnchorsAccountResource created on azure
// for more information of creating SpatialAnchorsAccountResource, please refer to the document of SpatialAnchorsAccountResource
string subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
string resourceGroupName = "MyResourceGroup";
string accountName = "MyAccount";
ResourceIdentifier spatialAnchorsAccountResourceId = SpatialAnchorsAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
SpatialAnchorsAccountResource spatialAnchorsAccount = client.GetSpatialAnchorsAccountResource(spatialAnchorsAccountResourceId);

// invoke the operation
MixedRealityAccountKeyRegenerateContent content = new MixedRealityAccountKeyRegenerateContent()
{
    Serial = MixedRealityAccountKeySerial.Primary,
};
MixedRealityAccountKeys result = await spatialAnchorsAccount.RegenerateKeysAsync(content);

Console.WriteLine($"Succeeded: {result}");
