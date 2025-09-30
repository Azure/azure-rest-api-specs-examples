using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.OracleDatabase.Models;
using Azure.ResourceManager.OracleDatabase;

// Generated from example definition: 2025-09-01/ExascaleDbStorageVaults_Get_MinimumSet_Gen.json
// this example is just showing the usage of "ExascaleDbStorageVault_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ExascaleDBStorageVaultResource created on azure
// for more information of creating ExascaleDBStorageVaultResource, please refer to the document of ExascaleDBStorageVaultResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rgopenapi";
string exascaleDbStorageVaultName = "storagevault1";
ResourceIdentifier exascaleDBStorageVaultResourceId = ExascaleDBStorageVaultResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, exascaleDbStorageVaultName);
ExascaleDBStorageVaultResource exascaleDBStorageVault = client.GetExascaleDBStorageVaultResource(exascaleDBStorageVaultResourceId);

// invoke the operation
ExascaleDBStorageVaultResource result = await exascaleDBStorageVault.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ExascaleDBStorageVaultData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
