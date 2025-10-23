using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Qumulo.Models;
using Azure.ResourceManager.Qumulo;

// Generated from example definition: 2024-06-19/FileSystems_Update_MaximumSet_Gen.json
// this example is just showing the usage of "FileSystemResource_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this QumuloFileSystemResource created on azure
// for more information of creating QumuloFileSystemResource, please refer to the document of QumuloFileSystemResource
string subscriptionId = "382E8C7A-AC80-4D70-8580-EFE99537B9B7";
string resourceGroupName = "rgQumulo";
string fileSystemName = "ahpixnvykleksjlr";
ResourceIdentifier qumuloFileSystemResourceId = QumuloFileSystemResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, fileSystemName);
QumuloFileSystemResource qumuloFileSystemResource = client.GetQumuloFileSystemResource(qumuloFileSystemResourceId);

// invoke the operation
QumuloFileSystemResourcePatch patch = new QumuloFileSystemResourcePatch
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
    ["key357"] = "ztkkvhfia"
    },
    Properties = new FileSystemResourceUpdateProperties
    {
        MarketplaceDetails = new MarketplaceDetails("fwtpz", "s")
        {
            MarketplaceSubscriptionId = "xaqtkloiyovmexqhn",
            PublisherId = "czxcfrwodazyaft",
            TermUnit = "cfwwczmygsimcyvoclcw",
        },
        UserDetailsEmail = "aqsnzyroo",
        DelegatedSubnetId = new ResourceIdentifier("bqaryqsjlackxphpmzffgoqsvm"),
    },
};
QumuloFileSystemResource result = await qumuloFileSystemResource.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
QumuloFileSystemResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
