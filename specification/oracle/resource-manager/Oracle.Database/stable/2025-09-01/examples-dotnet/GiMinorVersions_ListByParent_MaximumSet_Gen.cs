using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.OracleDatabase.Models;
using Azure.ResourceManager.OracleDatabase;

// Generated from example definition: 2025-09-01/GiMinorVersions_ListByParent_MaximumSet_Gen.json
// this example is just showing the usage of "GiMinorVersion_ListByParent" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this OracleGIVersionResource created on azure
// for more information of creating OracleGIVersionResource, please refer to the document of OracleGIVersionResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
AzureLocation location = new AzureLocation("eastus");
string giversionname = "name1";
ResourceIdentifier oracleGIVersionResourceId = OracleGIVersionResource.CreateResourceIdentifier(subscriptionId, location, giversionname);
OracleGIVersionResource oracleGIVersion = client.GetOracleGIVersionResource(oracleGIVersionResourceId);

// get the collection of this OracleGIMinorVersionResource
OracleGIMinorVersionCollection collection = oracleGIVersion.GetOracleGIMinorVersions();

// invoke the operation and iterate over the result
GIMinorVersionShapeFamily? shapeFamily = GIMinorVersionShapeFamily.Exadata;
string zone = "zone1";
await foreach (OracleGIMinorVersionResource item in collection.GetAllAsync(shapeFamily: shapeFamily, zone: zone))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    OracleGIMinorVersionData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
