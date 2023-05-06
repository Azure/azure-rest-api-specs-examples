using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Qumulo;
using Azure.ResourceManager.Qumulo.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/liftrqumulo/resource-manager/Qumulo.Storage/stable/2022-10-12/examples/FileSystems_Update_MaximumSet_Gen.json
// this example is just showing the usage of "FileSystems_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this QumuloFileSystemResource created on azure
// for more information of creating QumuloFileSystemResource, please refer to the document of QumuloFileSystemResource
string subscriptionId = "ulseeqylxb";
string resourceGroupName = "rgQumulo";
string fileSystemName = "nauwwbfoqehgbhdsmkewoboyxeqg";
ResourceIdentifier qumuloFileSystemResourceId = QumuloFileSystemResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, fileSystemName);
QumuloFileSystemResource qumuloFileSystemResource = client.GetQumuloFileSystemResource(qumuloFileSystemResourceId);

// invoke the operation
QumuloFileSystemResourcePatch patch = new QumuloFileSystemResourcePatch()
{
    Identity = new ManagedServiceIdentity("None")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("key4522")] = new UserAssignedIdentity(),
        },
    },
    Tags =
    {
    ["key7534"] = "jsgqvqbagquvxowbrkanyhzvo",
    },
    Properties = new FileSystemResourceUpdateProperties()
    {
        MarketplaceDetails = new MarketplaceDetails("x", "eiyhbmpwgezcmzrrfoiskuxlcvwojf", "wfmokfdjbwpjhz")
        {
            MarketplaceSubscriptionId = "ujrcqvxfnhxxheoth",
        },
        UserDetailsEmail = "aa",
        DelegatedSubnetId = new ResourceIdentifier("vjfirtaljehawmflyfianw"),
        ClusterLoginUri = new Uri("adabmuthwrbjshzfbo"),
        PrivateIPs =
        {
        "eugjqbaoucgjsopzfrq"
        },
    },
};
QumuloFileSystemResource result = await qumuloFileSystemResource.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
QumuloFileSystemResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
