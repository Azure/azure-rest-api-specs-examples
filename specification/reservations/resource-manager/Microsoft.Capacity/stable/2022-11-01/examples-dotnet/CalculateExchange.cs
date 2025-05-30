using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Reservations.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Reservations;

// Generated from example definition: specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-11-01/examples/CalculateExchange.json
// this example is just showing the usage of "CalculateExchange_Post" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
CalculateExchangeContent content = new CalculateExchangeContent
{
    Properties = new CalculateExchangeContentProperties
    {
        ReservationsToPurchase = {new ReservationPurchaseContent
        {
        SkuName = "Standard_B1ls",
        Location = new AzureLocation("westus"),
        ReservedResourceType = ReservedResourceType.VirtualMachines,
        BillingScopeId = new ResourceIdentifier("/subscriptions/ed3a1871-612d-abcd-a849-c2542a68be83"),
        Term = ReservationTerm.P1Y,
        BillingPlan = ReservationBillingPlan.Upfront,
        Quantity = 1,
        DisplayName = "testDisplayName",
        AppliedScopeType = AppliedScopeType.Shared,
        AppliedScopes = {},
        IsRenewEnabled = false,
        ReservedResourceInstanceFlexibility = InstanceFlexibility.On,
        }},
        SavingsPlansToPurchase = {new SavingsPlanPurchase
        {
        SkuName = "Compute_Savings_Plan",
        DisplayName = "ComputeSavingsPlan",
        BillingScopeId = new ResourceIdentifier("/subscriptions/10000000-0000-0000-0000-000000000000"),
        Term = SavingsPlanTerm.P1Y,
        AppliedScopeType = AppliedScopeType.Single,
        AppliedScopeProperties = new AppliedScopeProperties
        {
        ResourceGroupId = new ResourceIdentifier("/subscriptions/10000000-0000-0000-0000-000000000000/resourceGroups/testrg"),
        },
        Commitment = new BenefitsCommitment
        {
        Grain = BenefitsCommitmentGrain.Hourly,
        CurrencyCode = "USD",
        Amount = 15.23,
        },
        }},
        ReservationsToExchange = {new ReservationToReturn
        {
        ReservationId = new ResourceIdentifier("/providers/microsoft.capacity/reservationOrders/1f14354c-dc12-4c8d-8090-6f295a3a34aa/reservations/c8c926bd-fc5d-4e29-9d43-b68340ac23a6"),
        Quantity = 1,
        }},
    },
};
ArmOperation<CalculateExchangeResult> lro = await tenantResource.CalculateReservationExchangeAsync(WaitUntil.Completed, content);
CalculateExchangeResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
