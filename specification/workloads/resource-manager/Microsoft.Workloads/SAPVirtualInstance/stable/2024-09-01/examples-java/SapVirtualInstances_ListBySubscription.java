
/**
 * Samples for SapVirtualInstances List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapVirtualInstances_ListBySubscription.json
     */
    /**
     * Sample code: SAPVirtualInstances_ListBySubscription.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void sAPVirtualInstancesListBySubscription(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapVirtualInstances().list(com.azure.core.util.Context.NONE);
    }
}
