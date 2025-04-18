
/**
 * Samples for SapApplicationServerInstances List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapApplicationServerInstances_ListBySapVirtualInstance.json
     */
    /**
     * Sample code: SapApplicationServerInstances List By SAP Virtual Instance.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void sapApplicationServerInstancesListBySAPVirtualInstance(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapApplicationServerInstances().list("test-rg", "X00", com.azure.core.util.Context.NONE);
    }
}
