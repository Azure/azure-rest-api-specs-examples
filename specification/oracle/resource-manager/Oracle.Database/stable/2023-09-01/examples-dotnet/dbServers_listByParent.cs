using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.OracleDatabase;

// Generated from example definition: specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/dbServers_listByParent.json
// this example is just showing the usage of "DbServers_ListByCloudExadataInfrastructure" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CloudExadataInfrastructureResource created on azure
// for more information of creating CloudExadataInfrastructureResource, please refer to the document of CloudExadataInfrastructureResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg000";
string cloudexadatainfrastructurename = "infra1";
ResourceIdentifier cloudExadataInfrastructureResourceId = CloudExadataInfrastructureResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, cloudexadatainfrastructurename);
CloudExadataInfrastructureResource cloudExadataInfrastructure = client.GetCloudExadataInfrastructureResource(cloudExadataInfrastructureResourceId);

// get the collection of this OracleDBServerResource
OracleDBServerCollection collection = cloudExadataInfrastructure.GetOracleDBServers();

// invoke the operation and iterate over the result
await foreach (OracleDBServerResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    OracleDBServerData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
