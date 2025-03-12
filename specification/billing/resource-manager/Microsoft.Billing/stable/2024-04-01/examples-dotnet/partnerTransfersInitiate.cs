using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing.Models;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/partnerTransfersInitiate.json
// this example is just showing the usage of "PartnerTransfers_Initiate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PartnerTransferDetailResource created on azure
// for more information of creating PartnerTransferDetailResource, please refer to the document of PartnerTransferDetailResource
string billingAccountName = "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
string billingProfileName = "xxxx-xxxx-xxx-xxx";
string customerName = "11111111-1111-1111-1111-111111111111";
string transferName = "aabb123";
ResourceIdentifier partnerTransferDetailResourceId = PartnerTransferDetailResource.CreateResourceIdentifier(billingAccountName, billingProfileName, customerName, transferName);
PartnerTransferDetailResource partnerTransferDetail = client.GetPartnerTransferDetailResource(partnerTransferDetailResourceId);

// invoke the operation
PartnerTransferDetailCreateOrUpdateContent content = new PartnerTransferDetailCreateOrUpdateContent
{
    RecipientEmailId = "user@contoso.com",
};
ArmOperation<PartnerTransferDetailResource> lro = await partnerTransferDetail.UpdateAsync(WaitUntil.Completed, content);
PartnerTransferDetailResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PartnerTransferDetailData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
