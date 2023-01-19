using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ServiceLinker;
using Azure.ResourceManager.ServiceLinker.Models;

// Generated from example definition: specification/servicelinker/resource-manager/Microsoft.ServiceLinker/stable/2022-05-01/examples/PutLinkWithServiceEndpoint.json
// this example is just showing the usage of "Linker_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmResource created on azure
// for more information of creating ArmResource, please refer to the document of ArmResource

// get the collection of this LinkerResource
string resourceUri = "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app";
ResourceIdentifier scopeId = new ResourceIdentifier(string.Format("/{0}", resourceUri));
LinkerResourceCollection collection = client.GetLinkerResources(scopeId);

// invoke the operation
string linkerName = "linkName";
LinkerResourceData data = new LinkerResourceData()
{
    TargetService = new AzureResourceInfo()
    {
        Id = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DBforPostgreSQL/servers/test-pg/databases/test-db"),
    },
    AuthInfo = new SecretAuthInfo()
    {
        Name = "name",
        SecretInfo = new KeyVaultSecretUriSecretInfo()
        {
            Value = "https://vault-name.vault.azure.net/secrets/secret-name/00000000000000000000000000000000",
        },
    },
    SolutionType = VnetSolutionType.ServiceEndpoint,
};
ArmOperation<LinkerResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, linkerName, data);
LinkerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
LinkerResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
