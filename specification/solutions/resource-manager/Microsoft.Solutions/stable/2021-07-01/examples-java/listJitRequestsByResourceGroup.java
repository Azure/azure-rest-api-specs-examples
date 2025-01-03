
/**
 * Samples for JitRequests ListBySubscription.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/
     * listJitRequestsByResourceGroup.json
     */
    /**
     * Sample code: Lists all JIT requests within the subscription.
     * 
     * @param manager Entry point to ApplicationManager.
     */
    public static void listsAllJITRequestsWithinTheSubscription(
        com.azure.resourcemanager.managedapplications.ApplicationManager manager) {
        manager.jitRequests().listBySubscriptionWithResponse(com.azure.core.util.Context.NONE);
    }
}
