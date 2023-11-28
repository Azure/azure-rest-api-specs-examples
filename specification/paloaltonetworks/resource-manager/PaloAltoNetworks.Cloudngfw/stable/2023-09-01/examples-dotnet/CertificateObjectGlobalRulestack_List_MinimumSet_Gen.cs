using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw.Models;

// Generated from example definition: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/CertificateObjectGlobalRulestack_List_MinimumSet_Gen.json
// this example is just showing the usage of "CertificateObjectGlobalRulestack_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this GlobalRulestackResource created on azure
// for more information of creating GlobalRulestackResource, please refer to the document of GlobalRulestackResource
string globalRulestackName = "praval";
ResourceIdentifier globalRulestackResourceId = GlobalRulestackResource.CreateResourceIdentifier(globalRulestackName);
GlobalRulestackResource globalRulestack = client.GetGlobalRulestackResource(globalRulestackResourceId);

// get the collection of this GlobalRulestackCertificateObjectResource
GlobalRulestackCertificateObjectCollection collection = globalRulestack.GetGlobalRulestackCertificateObjects();

// invoke the operation and iterate over the result
await foreach (GlobalRulestackCertificateObjectResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    GlobalRulestackCertificateObjectData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
