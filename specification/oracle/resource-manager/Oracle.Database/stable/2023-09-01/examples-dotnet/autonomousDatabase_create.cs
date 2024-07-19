using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.OracleDatabase.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.OracleDatabase;

// Generated from example definition: specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/autonomousDatabase_create.json
// this example is just showing the usage of "AutonomousDatabases_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg000";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this AutonomousDatabaseResource
AutonomousDatabaseCollection collection = resourceGroupResource.GetAutonomousDatabases();

// invoke the operation
string autonomousdatabasename = "databasedb1";
AutonomousDatabaseData data = new AutonomousDatabaseData(new AzureLocation("eastus"))
{
    Properties = new AutonomousDatabaseProperties()
    {
        AdminPassword = "********",
        CharacterSet = "AL32UTF8",
        ComputeCount = 2,
        ComputeModel = AutonomousDatabaseComputeModel.Ecpu,
        DataStorageSizeInTbs = 1,
        DBVersion = "18.4.0.0",
        DisplayName = "example_autonomous_databasedb1",
        NcharacterSet = "AL16UTF16",
        SubnetId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
        VnetId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Microsoft.Network/virtualNetworks/vnet1"),
    },
    Tags =
    {
    ["tagK1"] = "tagV1",
    },
};
ArmOperation<AutonomousDatabaseResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, autonomousdatabasename, data);
AutonomousDatabaseResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AutonomousDatabaseData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
