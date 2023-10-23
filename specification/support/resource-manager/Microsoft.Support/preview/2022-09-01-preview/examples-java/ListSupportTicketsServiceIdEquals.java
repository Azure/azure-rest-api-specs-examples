/** Samples for SupportTicketsNoSubscription List. */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/ListSupportTicketsServiceIdEquals.json
     */
    /**
     * Sample code: List support tickets with a certain service id for a subscription.
     *
     * @param manager Entry point to SupportManager.
     */
    public static void listSupportTicketsWithACertainServiceIdForASubscription(
        com.azure.resourcemanager.support.SupportManager manager) {
        manager
            .supportTicketsNoSubscriptions()
            .list(null, "ServiceId eq 'vm_windows_service_guid'", com.azure.core.util.Context.NONE);
    }
}
