/** Samples for SupportTicketsNoSubscription List. */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/ListSupportTicketsCreatedOnOrAfterAndInOpenState.json
     */
    /**
     * Sample code: List support tickets created on or after a certain date and in open state for a subscription.
     *
     * @param manager Entry point to SupportManager.
     */
    public static void listSupportTicketsCreatedOnOrAfterACertainDateAndInOpenStateForASubscription(
        com.azure.resourcemanager.support.SupportManager manager) {
        manager
            .supportTicketsNoSubscriptions()
            .list(null, "createdDate ge 2020-03-10T22:08:51Z and status eq 'Open'", com.azure.core.util.Context.NONE);
    }
}
