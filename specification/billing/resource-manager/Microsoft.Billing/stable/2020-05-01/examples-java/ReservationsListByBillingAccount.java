/** Samples for Reservations ListByBillingAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/ReservationsListByBillingAccount.json
     */
    /**
     * Sample code: ReservationsListByBillingAccount.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void reservationsListByBillingAccount(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .reservations()
            .listByBillingAccount(
                "{billingAccountName}",
                "properties/reservedResourceType eq 'VirtualMachines'",
                "properties/userFriendlyAppliedScopeType asc",
                "true",
                "Succeeded",
                com.azure.core.util.Context.NONE);
    }
}
