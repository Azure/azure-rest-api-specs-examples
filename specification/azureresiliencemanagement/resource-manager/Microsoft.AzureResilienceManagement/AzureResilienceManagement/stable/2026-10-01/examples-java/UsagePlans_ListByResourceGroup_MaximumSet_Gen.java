
/**
 * Samples for UsagePlans ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/UsagePlans_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: UsagePlans_ListByResourceGroup_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void usagePlansListByResourceGroupMaximumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.usagePlans().listByResourceGroup("MyResourceGroup", com.azure.core.util.Context.NONE);
    }
}
