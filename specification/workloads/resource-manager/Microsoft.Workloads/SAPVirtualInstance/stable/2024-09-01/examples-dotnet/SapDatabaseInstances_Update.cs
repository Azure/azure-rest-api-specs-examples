using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.WorkloadsSapVirtualInstance.Models;
using Azure.ResourceManager.WorkloadsSapVirtualInstance;

// Generated from example definition: 2024-09-01/SapDatabaseInstances_Update.json
// this example is just showing the usage of "SAPDatabaseInstance_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SapDatabaseInstanceResource created on azure
// for more information of creating SapDatabaseInstanceResource, please refer to the document of SapDatabaseInstanceResource
string subscriptionId = "6d875e77-e412-4d7d-9af4-8895278b4443";
string resourceGroupName = "test-rg";
string sapVirtualInstanceName = "X00";
string databaseInstanceName = "databaseServer";
ResourceIdentifier sapDatabaseInstanceResourceId = SapDatabaseInstanceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, sapVirtualInstanceName, databaseInstanceName);
SapDatabaseInstanceResource sapDatabaseInstance = client.GetSapDatabaseInstanceResource(sapDatabaseInstanceResourceId);

// invoke the operation
SapDatabaseInstancePatch patch = new SapDatabaseInstancePatch
{
    Tags =
    {
    ["key1"] = "value1"
    },
};
SapDatabaseInstanceResource result = await sapDatabaseInstance.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SapDatabaseInstanceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
