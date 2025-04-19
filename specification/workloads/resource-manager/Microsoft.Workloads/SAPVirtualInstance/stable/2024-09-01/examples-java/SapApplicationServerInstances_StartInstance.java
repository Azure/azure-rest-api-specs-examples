
import com.azure.resourcemanager.workloadssapvirtualinstance.models.StartRequest;

/**
 * Samples for SapApplicationServerInstances Start.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapApplicationServerInstances_StartInstance.json
     */
    /**
     * Sample code: Start the SAP Application Server Instance.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void startTheSAPApplicationServerInstance(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapApplicationServerInstances().start("test-rg", "X00", "app01", new StartRequest(),
            com.azure.core.util.Context.NONE);
    }
}
