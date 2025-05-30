using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing.Models;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/recipientTransfersAccept.json
// this example is just showing the usage of "RecipientTransfers_Accept" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RecipientTransferDetailResource created on azure
// for more information of creating RecipientTransferDetailResource, please refer to the document of RecipientTransferDetailResource
string transferName = "aabb123";
ResourceIdentifier recipientTransferDetailResourceId = RecipientTransferDetailResource.CreateResourceIdentifier(transferName);
RecipientTransferDetailResource recipientTransferDetail = client.GetRecipientTransferDetailResource(recipientTransferDetailResourceId);

// invoke the operation
AcceptTransferContent content = new AcceptTransferContent
{
    ProductDetails = {new BillingProductDetails
    {
    ProductType = BillingProductType.AzureSubscription,
    ProductId = "subscriptionId",
    }, new BillingProductDetails
    {
    ProductType = BillingProductType.AzureReservation,
    ProductId = "reservedInstanceId",
    }},
};
RecipientTransferDetailResource result = await recipientTransferDetail.AcceptAsync(content);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RecipientTransferDetailData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
