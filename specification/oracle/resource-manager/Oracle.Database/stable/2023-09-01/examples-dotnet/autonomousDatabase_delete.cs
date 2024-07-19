using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.OracleDatabase.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.OracleDatabase;

// Generated from example definition: specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/autonomousDatabase_delete.json
// this example is just showing the usage of "AutonomousDatabases_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AutonomousDatabaseResource created on azure
// for more information of creating AutonomousDatabaseResource, please refer to the document of AutonomousDatabaseResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg000";
string autonomousdatabasename = "databasedb1";
ResourceIdentifier autonomousDatabaseResourceId = AutonomousDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, autonomousdatabasename);
AutonomousDatabaseResource autonomousDatabase = client.GetAutonomousDatabaseResource(autonomousDatabaseResourceId);

// invoke the operation
await autonomousDatabase.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
