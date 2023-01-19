/** Samples for SupportTickets List. */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/ListSupportTicketsProblemClassificationIdEqualsForSubscription.json
     */
    /**
     * Sample code: List support tickets with a certain problem classification id for a subscription.
     *
     * @param manager Entry point to SupportManager.
     */
    public static void listSupportTicketsWithACertainProblemClassificationIdForASubscription(
        com.azure.resourcemanager.support.SupportManager manager) {
        manager
            .supportTickets()
            .list(
                null,
                "ProblemClassificationId eq 'compute_vm_problemClassification_guid'",
                com.azure.core.util.Context.NONE);
    }
}
