using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RecoveryServicesBackup.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.RecoveryServicesBackup;

// Generated from example definition: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/TieringCost/FetchTieringCostForPolicy.json
// this example is just showing the usage of "FetchTieringCost_Post" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "netsdktestrg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// invoke the operation
string vaultName = "testVault";
FetchTieringCostInfoContent content = new FetchTieringCostSavingsInfoForPolicyContent(RecoveryPointTierType.HardenedRP, RecoveryPointTierType.ArchivedRP, "monthly");
ArmOperation<TieringCostInfo> lro = await resourceGroupResource.PostFetchTieringCostAsync(WaitUntil.Completed, vaultName, content);
TieringCostInfo result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
