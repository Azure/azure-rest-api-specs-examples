using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw.Models;

// Generated from example definition: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/preview/2022-08-29-preview/examples/CertificateObjectLocalRulestack_Delete_MinimumSet_Gen.json
// this example is just showing the usage of "CertificateObjectLocalRulestack_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CertificateObjectLocalRulestackResource created on azure
// for more information of creating CertificateObjectLocalRulestackResource, please refer to the document of CertificateObjectLocalRulestackResource
string subscriptionId = "2bf4a339-294d-4c25-b0b2-ef649e9f5c27";
string resourceGroupName = "rgopenapi";
string localRulestackName = "lrs1";
string name = "armid1";
ResourceIdentifier certificateObjectLocalRulestackResourceId = CertificateObjectLocalRulestackResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, localRulestackName, name);
CertificateObjectLocalRulestackResource certificateObjectLocalRulestackResource = client.GetCertificateObjectLocalRulestackResource(certificateObjectLocalRulestackResourceId);

// invoke the operation
await certificateObjectLocalRulestackResource.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
