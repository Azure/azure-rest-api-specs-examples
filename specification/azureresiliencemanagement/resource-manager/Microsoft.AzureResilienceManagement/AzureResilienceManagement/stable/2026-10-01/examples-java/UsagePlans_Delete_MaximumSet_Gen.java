
/**
 * Samples for UsagePlans Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/UsagePlans_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: UsagePlans_Delete_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void
        usagePlansDeleteMaximumSet(com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.usagePlans().delete("MyResourceGroup", "myUsagePlan", com.azure.core.util.Context.NONE);
    }
}
