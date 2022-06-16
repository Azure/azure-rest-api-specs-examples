import com.azure.core.util.Context;

/** Samples for Reservations ListByBillingProfile. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/ReservationsListByBillingProfile.json
     */
    /**
     * Sample code: ReservationsListByBillingProfile.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void reservationsListByBillingProfile(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .reservations()
            .listByBillingProfile(
                "{billingAccountName}",
                "{billingProfileName}",
                "properties/reservedResourceType eq 'VirtualMachines'",
                "properties/userFriendlyAppliedScopeType asc",
                "true",
                "Succeeded",
                Context.NONE);
    }
}
