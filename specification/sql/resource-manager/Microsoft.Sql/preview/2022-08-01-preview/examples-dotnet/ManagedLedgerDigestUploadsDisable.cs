using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;
using Azure.ResourceManager.Sql.Models;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/ManagedLedgerDigestUploadsDisable.json
// this example is just showing the usage of "ManagedLedgerDigestUploads_Disable" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedLedgerDigestUploadResource created on azure
// for more information of creating ManagedLedgerDigestUploadResource, please refer to the document of ManagedLedgerDigestUploadResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "ledgertestrg";
string managedInstanceName = "ledgertestserver";
string databaseName = "testdb";
ManagedLedgerDigestUploadsName ledgerDigestUploads = ManagedLedgerDigestUploadsName.Current;
ResourceIdentifier managedLedgerDigestUploadResourceId = ManagedLedgerDigestUploadResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, managedInstanceName, databaseName, ledgerDigestUploads);
ManagedLedgerDigestUploadResource managedLedgerDigestUpload = client.GetManagedLedgerDigestUploadResource(managedLedgerDigestUploadResourceId);

// invoke the operation
ArmOperation<ManagedLedgerDigestUploadResource> lro = await managedLedgerDigestUpload.DisableAsync(WaitUntil.Completed);
ManagedLedgerDigestUploadResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagedLedgerDigestUploadData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
