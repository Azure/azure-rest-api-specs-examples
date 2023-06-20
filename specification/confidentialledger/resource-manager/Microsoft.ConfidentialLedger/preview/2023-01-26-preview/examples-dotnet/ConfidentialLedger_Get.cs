using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ConfidentialLedger;
using Azure.ResourceManager.ConfidentialLedger.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2023-01-26-preview/examples/ConfidentialLedger_Get.json
// this example is just showing the usage of "Ledger_Get" operation, for the dependent resources, they will have to be created separately.

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
ConfidentialLedgerResource result = await confidentialLedger.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ConfidentialLedgerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
