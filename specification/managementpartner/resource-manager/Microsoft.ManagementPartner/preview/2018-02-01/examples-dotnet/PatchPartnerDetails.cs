using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagementPartner;

// Generated from example definition: specification/managementpartner/resource-manager/Microsoft.ManagementPartner/preview/2018-02-01/examples/PatchPartnerDetails.json
// this example is just showing the usage of "Partner_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PartnerResponseResource created on azure
// for more information of creating PartnerResponseResource, please refer to the document of PartnerResponseResource
string partnerId = "123456";
ResourceIdentifier partnerResponseResourceId = PartnerResponseResource.CreateResourceIdentifier(partnerId);
PartnerResponseResource partnerResponse = client.GetPartnerResponseResource(partnerResponseResourceId);

// invoke the operation
PartnerResponseResource result = await partnerResponse.UpdateAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PartnerResponseData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
