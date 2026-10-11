
/**
 * Samples for UsagePlans List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/UsagePlans_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: UsagePlans_ListBySubscription_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void usagePlansListBySubscriptionMaximumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.usagePlans().list(com.azure.core.util.Context.NONE);
    }
}
