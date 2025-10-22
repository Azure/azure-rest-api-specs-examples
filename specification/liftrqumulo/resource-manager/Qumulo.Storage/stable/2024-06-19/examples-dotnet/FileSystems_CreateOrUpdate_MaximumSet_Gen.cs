using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Qumulo;

// Generated from example definition: 2024-06-19/FileSystems_CreateOrUpdate_MaximumSet_Gen.json
// this example is just showing the usage of "FileSystemResource_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "382E8C7A-AC80-4D70-8580-EFE99537B9B7";
string resourceGroupName = "rgQumulo";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this QumuloFileSystemResource
QumuloFileSystemResourceCollection collection = resourceGroupResource.GetQumuloFileSystemResources();

// invoke the operation
string fileSystemName = "hfcmtgaes";
QumuloFileSystemResourceData data = new QumuloFileSystemResourceData(new AzureLocation("pnb"))
{
    Identity = new ManagedServiceIdentity("None")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("key7679")] = new UserAssignedIdentity()
        },
    },
    Tags =
    {
    ["key7090"] = "rurrdiaqp"
    },
};
ArmOperation<QumuloFileSystemResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, fileSystemName, data);
QumuloFileSystemResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
QumuloFileSystemResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
