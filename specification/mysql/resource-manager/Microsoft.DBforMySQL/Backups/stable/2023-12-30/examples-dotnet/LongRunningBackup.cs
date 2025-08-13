using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MySql.FlexibleServers;

// Generated from example definition: specification/mysql/resource-manager/Microsoft.DBforMySQL/Backups/stable/2023-12-30/examples/LongRunningBackup.json
// this example is just showing the usage of "LongRunningBackup_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MySqlFlexibleServerBackupV2Resource created on azure
// for more information of creating MySqlFlexibleServerBackupV2Resource, please refer to the document of MySqlFlexibleServerBackupV2Resource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "TestGroup";
string serverName = "mysqltestserver";
string backupName = "testback";
ResourceIdentifier mySqlFlexibleServerBackupV2ResourceId = MySqlFlexibleServerBackupV2Resource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, backupName);
MySqlFlexibleServerBackupV2Resource mySqlFlexibleServerBackupV2 = client.GetMySqlFlexibleServerBackupV2Resource(mySqlFlexibleServerBackupV2ResourceId);

// invoke the operation
MySqlFlexibleServerBackupV2Data data = new MySqlFlexibleServerBackupV2Data();
ArmOperation<MySqlFlexibleServerBackupV2Resource> lro = await mySqlFlexibleServerBackupV2.UpdateAsync(WaitUntil.Completed, data);
MySqlFlexibleServerBackupV2Resource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MySqlFlexibleServerBackupV2Data resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
