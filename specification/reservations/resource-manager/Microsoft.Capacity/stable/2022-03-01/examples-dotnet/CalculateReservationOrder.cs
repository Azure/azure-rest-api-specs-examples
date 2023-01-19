using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Reservations;
using Azure.ResourceManager.Reservations.Models;

// Generated from example definition: specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-03-01/examples/CalculateReservationOrder.json
// this example is just showing the usage of "ReservationOrder_Calculate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TenantResource created on azure
// for more information of creating TenantResource, please refer to the document of TenantResource
var tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
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
    ReservedResourceInstanceFlexibility = InstanceFlexibility.On,
};
CalculatePriceResult result = await tenantResource.CalculateReservationOrderAsync(content);

Console.WriteLine($"Succeeded: {result}");
