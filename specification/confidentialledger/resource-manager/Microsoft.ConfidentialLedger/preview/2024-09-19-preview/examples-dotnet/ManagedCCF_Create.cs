using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ConfidentialLedger.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ConfidentialLedger;

// Generated from example definition: specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2024-09-19-preview/examples/ManagedCCF_Create.json
// this example is just showing the usage of "ManagedCCF_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "0000000-0000-0000-0000-000000000001";
string resourceGroupName = "DummyResourceGroupName";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ManagedCcfResource
ManagedCcfCollection collection = resourceGroupResource.GetManagedCcfs();

// invoke the operation
string appName = "DummyMccfAppName";
ManagedCcfData data = new ManagedCcfData(new AzureLocation("EastUS"))
{
    Properties = new ManagedCcfProperties
    {
        MemberIdentityCertificates = {new ConfidentialLedgerMemberIdentityCertificate
        {
        Certificate = "-----BEGIN CERTIFICATE-----MIIBsjCCATigAwIBAgIUZWIbyG79TniQLd2UxJuU74tqrKcwCgYIKoZIzj0EAwMwEDEOMAwGA1UEAwwFdXNlcjAwHhcNMjEwMzE2MTgwNjExWhcNMjIwMzE2MTgwNjExWjAQMQ4wDAYDVQQDDAV1c2VyMDB2MBAGByqGSM49AgEGBSuBBAAiA2IABBiWSo/j8EFit7aUMm5lF+lUmCu+IgfnpFD+7QMgLKtxRJ3aGSqgS/GpqcYVGddnODtSarNE/HyGKUFUolLPQ5ybHcouUk0kyfA7XMeSoUA4lBz63Wha8wmXo+NdBRo39qNTMFEwHQYDVR0OBBYEFPtuhrwgGjDFHeUUT4nGsXaZn69KMB8GA1UdIwQYMBaAFPtuhrwgGjDFHeUUT4nGsXaZn69KMA8GA1UdEwEB/wQFMAMBAf8wCgYIKoZIzj0EAwMDaAAwZQIxAOnozm2CyqRwSSQLls5r+mUHRGRyXHXwYtM4Dcst/VEZdmS9fqvHRCHbjUlO/+HNfgIwMWZ4FmsjD3wnPxONOm9YdVn/PRD7SsPRPbOjwBiE4EBGaHDsLjYAGDSGi7NJnSkA-----END CERTIFICATE-----",
        Encryptionkey = "ledgerencryptionkey",
        Tags = BinaryData.FromObjectAsJson(new
        {
        additionalProps1 = "additional properties",
        }),
        }},
        DeploymentType = new ConfidentialLedgerDeploymentType
        {
            LanguageRuntime = ConfidentialLedgerLanguageRuntime.CPP,
            AppSourceUri = new Uri("https://myaccount.blob.core.windows.net/storage/mccfsource?sv=2022-02-11%st=2022-03-11"),
        },
        NodeCount = 5,
    },
    Tags =
    {
    ["additionalProps1"] = "additional properties"
    },
};
ArmOperation<ManagedCcfResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, appName, data);
ManagedCcfResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagedCcfData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
