using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Reservations;
using Azure.ResourceManager.Reservations.Models;

// Generated from example definition: specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-11-01/examples/PurchaseReservationOrder.json
// this example is just showing the usage of "ReservationOrder_Purchase" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TenantResource created on azure
// for more information of creating TenantResource, please refer to the document of TenantResource
var tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// get the collection of this ReservationOrderResource
ReservationOrderCollection collection = tenantResource.GetReservationOrders();

// invoke the operation
Guid reservationOrderId = Guid.Parse("a075419f-44cc-497f-b68a-14ee811d48b9");
ReservationPurchaseContent content = new ReservationPurchaseContent()
{
    SkuName = "standard_D1",
    Location = new AzureLocation("westus"),
    ReservedResourceType = ReservedResourceType.VirtualMachines,
    BillingScopeId = new ResourceIdentifier("/subscriptions/ed3a1871-612d-abcd-a849-c2542a68be83"),
    Term = ReservationTerm.P1Y,
    BillingPlan = ReservationBillingPlan.Monthly,
    Quantity = 1,
    DisplayName = "TestReservationOrder",
    AppliedScopeType = AppliedScopeType.Shared,
    AppliedScopes =
    {
    },
    IsRenewEnabled = false,
    ReservedResourceInstanceFlexibility = InstanceFlexibility.On,
};
ArmOperation<ReservationOrderResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, reservationOrderId, content);
ReservationOrderResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ReservationOrderData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
