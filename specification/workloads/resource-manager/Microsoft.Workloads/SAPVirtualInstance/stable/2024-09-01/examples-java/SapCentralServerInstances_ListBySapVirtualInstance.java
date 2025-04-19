
/**
 * Samples for SapCentralServerInstances List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapCentralServerInstances_ListBySapVirtualInstance.json
     */
    /**
     * Sample code: SAPCentralInstances List by SAP virtual instance.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void sAPCentralInstancesListBySAPVirtualInstance(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapCentralServerInstances().list("test-rg", "X00", com.azure.core.util.Context.NONE);
    }
}
