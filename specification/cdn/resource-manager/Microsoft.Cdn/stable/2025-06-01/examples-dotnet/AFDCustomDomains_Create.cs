using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Cdn.Models;
using Azure.ResourceManager.Cdn;

// Generated from example definition: specification/cdn/resource-manager/Microsoft.Cdn/stable/2025-06-01/examples/AFDCustomDomains_Create.json
// this example is just showing the usage of "FrontDoorCustomDomains_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ProfileResource created on azure
// for more information of creating ProfileResource, please refer to the document of ProfileResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "RG";
string profileName = "profile1";
ResourceIdentifier profileResourceId = ProfileResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, profileName);
ProfileResource profile = client.GetProfileResource(profileResourceId);

// get the collection of this FrontDoorCustomDomainResource
FrontDoorCustomDomainCollection collection = profile.GetFrontDoorCustomDomains();

// invoke the operation
string customDomainName = "domain1";
FrontDoorCustomDomainData data = new FrontDoorCustomDomainData
{
    TlsSettings = new FrontDoorCustomDomainHttpsContent(FrontDoorCertificateType.ManagedCertificate)
    {
        CipherSuiteSetType = AfdCipherSuiteSetType.Customized,
        MinimumTlsVersion = FrontDoorMinimumTlsVersion.Tls1_2,
        CustomizedCipherSuiteSet = new FrontDoorCustomDomainHttpsCustomizedCipherSuiteSet
        {
            CipherSuiteSetForTls12 = { AfdCustomizedCipherSuiteForTls12.Ecdhe_Rsa_Aes128_Gcm_Sha256 },
            CipherSuiteSetForTls13 = { AfdCustomizedCipherSuiteForTls13.Tls_Aes_128_Gcm_Sha256, AfdCustomizedCipherSuiteForTls13.Tls_Aes_256_Gcm_Sha384 },
        },
    },
    DnsZoneId = new ResourceIdentifier(""),
    HostName = "www.someDomain.net",
};
ArmOperation<FrontDoorCustomDomainResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, customDomainName, data);
FrontDoorCustomDomainResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FrontDoorCustomDomainData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
