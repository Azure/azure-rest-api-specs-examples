
import com.azure.resourcemanager.reservations.models.AppliedScopeProperties;
import com.azure.resourcemanager.reservations.models.AppliedScopeType;
import com.azure.resourcemanager.reservations.models.CalculateExchangeRequest;
import com.azure.resourcemanager.reservations.models.CalculateExchangeRequestProperties;
import com.azure.resourcemanager.reservations.models.Commitment;
import com.azure.resourcemanager.reservations.models.CommitmentGrain;
import com.azure.resourcemanager.reservations.models.InstanceFlexibility;
import com.azure.resourcemanager.reservations.models.PurchaseRequest;
import com.azure.resourcemanager.reservations.models.PurchaseRequestPropertiesReservedResourceProperties;
import com.azure.resourcemanager.reservations.models.ReservationBillingPlan;
import com.azure.resourcemanager.reservations.models.ReservationTerm;
import com.azure.resourcemanager.reservations.models.ReservationToReturn;
import com.azure.resourcemanager.reservations.models.ReservedResourceType;
import com.azure.resourcemanager.reservations.models.SavingsPlanPurchaseRequest;
import com.azure.resourcemanager.reservations.models.SavingsPlanTerm;
import com.azure.resourcemanager.reservations.models.SkuName;
import java.util.Arrays;

/**
 * Samples for CalculateExchange Post.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-11-01/examples/CalculateExchange.json
     */
    /**
     * Sample code: CalculateExchange.
     * 
     * @param manager Entry point to ReservationsManager.
     */
    public static void calculateExchange(com.azure.resourcemanager.reservations.ReservationsManager manager) {
        manager.calculateExchanges()
            .post(new CalculateExchangeRequest().withProperties(new CalculateExchangeRequestProperties()
                .withReservationsToPurchase(Arrays.asList(new PurchaseRequest()
                    .withSku(new SkuName().withName("Standard_B1ls")).withLocation("westus")
                    .withReservedResourceType(ReservedResourceType.VIRTUAL_MACHINES)
                    .withBillingScopeId("/subscriptions/ed3a1871-612d-abcd-a849-c2542a68be83")
                    .withTerm(ReservationTerm.P1Y).withBillingPlan(ReservationBillingPlan.UPFRONT).withQuantity(1)
                    .withDisplayName("testDisplayName").withAppliedScopeType(AppliedScopeType.SHARED).withRenew(false)
                    .withReservedResourceProperties(new PurchaseRequestPropertiesReservedResourceProperties()
                        .withInstanceFlexibility(InstanceFlexibility.ON))))
                .withSavingsPlansToPurchase(Arrays.asList(new SavingsPlanPurchaseRequest()
                    .withSku(new SkuName().withName("Compute_Savings_Plan")).withDisplayName("ComputeSavingsPlan")
                    .withBillingScopeId("/subscriptions/10000000-0000-0000-0000-000000000000")
                    .withTerm(SavingsPlanTerm.P1Y).withAppliedScopeType(AppliedScopeType.SINGLE)
                    .withAppliedScopeProperties(new AppliedScopeProperties().withResourceGroupId(
                        "/subscriptions/10000000-0000-0000-0000-000000000000/resourceGroups/testrg"))
                    .withCommitment(new Commitment().withCurrencyCode("fakeTokenPlaceholder").withAmount(15.23D)
                        .withGrain(CommitmentGrain.HOURLY))))
                .withReservationsToExchange(Arrays.asList(new ReservationToReturn().withReservationId(
                    "/providers/microsoft.capacity/reservationOrders/1f14354c-dc12-4c8d-8090-6f295a3a34aa/reservations/c8c926bd-fc5d-4e29-9d43-b68340ac23a6")
                    .withQuantity(1)))),
                com.azure.core.util.Context.NONE);
    }
}
