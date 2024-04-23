
/**
 * Samples for SupportTicketsNoSubscription List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/
     * ListSupportTicketsInUpdatingState.json
     */
    /**
     * Sample code: List support tickets in updating state.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void listSupportTicketsInUpdatingState(com.azure.resourcemanager.support.SupportManager manager) {
        manager.supportTicketsNoSubscriptions().list(null, "status eq 'Updating'", com.azure.core.util.Context.NONE);
    }

    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/
     * ListSupportTicketsServiceIdEquals.json
     */
    /**
     * Sample code: List support tickets with a certain service id.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void
        listSupportTicketsWithACertainServiceId(com.azure.resourcemanager.support.SupportManager manager) {
        manager.supportTicketsNoSubscriptions().list(null, "ServiceId eq 'vm_windows_service_guid'",
            com.azure.core.util.Context.NONE);
    }

    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/
     * ListSupportTicketsProblemClassificationIdEquals.json
     */
    /**
     * Sample code: List support tickets with a certain problem classification id.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void listSupportTicketsWithACertainProblemClassificationId(
        com.azure.resourcemanager.support.SupportManager manager) {
        manager.supportTicketsNoSubscriptions().list(null,
            "ProblemClassificationId eq 'compute_vm_problemClassification_guid'", com.azure.core.util.Context.NONE);
    }
}
