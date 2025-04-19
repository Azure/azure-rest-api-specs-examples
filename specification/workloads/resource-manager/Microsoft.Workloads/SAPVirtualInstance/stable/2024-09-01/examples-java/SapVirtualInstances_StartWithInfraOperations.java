
import com.azure.resourcemanager.workloadssapvirtualinstance.models.StartRequest;

/**
 * Samples for SapVirtualInstances Start.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapVirtualInstances_StartWithInfraOperations.json
     */
    /**
     * Sample code: SAPVirtualInstances_Start_WithInfraOperations.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void sAPVirtualInstancesStartWithInfraOperations(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapVirtualInstances().start("test-rg", "X00", new StartRequest().withStartVm(true),
            com.azure.core.util.Context.NONE);
    }
}
