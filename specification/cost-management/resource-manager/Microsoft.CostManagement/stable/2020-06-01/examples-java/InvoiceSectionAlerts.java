/** Samples for Alerts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2020-06-01/examples/InvoiceSectionAlerts.json
     */
    /**
     * Sample code: InvoiceSectionAlerts.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void invoiceSectionAlerts(com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .alerts()
            .listWithResponse(
                "providers/Microsoft.Billing/billingAccounts/12345:6789/billingProfiles/13579/invoiceSections/9876",
                com.azure.core.util.Context.NONE);
    }
}
