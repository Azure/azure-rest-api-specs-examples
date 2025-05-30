using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw.Models;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw;

// Generated from example definition: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/CertificateObjectLocalRulestack_CreateOrUpdate_MinimumSet_Gen.json
// this example is just showing the usage of "CertificateObjectLocalRulestack_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this LocalRulestackCertificateObjectResource created on azure
// for more information of creating LocalRulestackCertificateObjectResource, please refer to the document of LocalRulestackCertificateObjectResource
string subscriptionId = "2bf4a339-294d-4c25-b0b2-ef649e9f5c27";
string resourceGroupName = "rgopenapi";
string localRulestackName = "lrs1";
string name = "armid1";
ResourceIdentifier localRulestackCertificateObjectResourceId = LocalRulestackCertificateObjectResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, localRulestackName, name);
LocalRulestackCertificateObjectResource localRulestackCertificateObject = client.GetLocalRulestackCertificateObjectResource(localRulestackCertificateObjectResourceId);

// invoke the operation
LocalRulestackCertificateObjectData data = new LocalRulestackCertificateObjectData(FirewallBooleanType.True);
ArmOperation<LocalRulestackCertificateObjectResource> lro = await localRulestackCertificateObject.UpdateAsync(WaitUntil.Completed, data);
LocalRulestackCertificateObjectResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
LocalRulestackCertificateObjectData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
