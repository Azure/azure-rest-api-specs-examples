using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ConfidentialLedger.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ConfidentialLedger;

// Generated from example definition: specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2023-06-28-preview/examples/ConfidentialLedger_Restore.json
// this example is just showing the usage of "Ledger_Restore" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ConfidentialLedgerResource created on azure
// for more information of creating ConfidentialLedgerResource, please refer to the document of ConfidentialLedgerResource
string subscriptionId = "0000000-0000-0000-0000-000000000001";
string resourceGroupName = "DummyResourceGroupName";
string ledgerName = "DummyLedgerName";
ResourceIdentifier confidentialLedgerResourceId = ConfidentialLedgerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, ledgerName);
ConfidentialLedgerResource confidentialLedger = client.GetConfidentialLedgerResource(confidentialLedgerResourceId);

// invoke the operation
ConfidentialLedgerRestoreContent content = new ConfidentialLedgerRestoreContent("DummyFileShareName", "EastUS", new Uri("DummySASUri"));
ArmOperation<ConfidentialLedgerRestoreResult> lro = await confidentialLedger.RestoreAsync(WaitUntil.Completed, content);
ConfidentialLedgerRestoreResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
