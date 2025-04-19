
import com.azure.resourcemanager.workloadssapvirtualinstance.models.StopRequest;

/**
 * Samples for SapApplicationServerInstances Stop.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapApplicationServerInstances_StopInstanceInfrastructure.json
     */
    /**
     * Sample code: Stop the SAP Application Server Instance and it's infrastructure.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void stopTheSAPApplicationServerInstanceAndItSInfrastructure(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapApplicationServerInstances().stop("test-rg", "X00", "app01",
            new StopRequest().withSoftStopTimeoutSeconds(0L).withDeallocateVm(true), com.azure.core.util.Context.NONE);
    }
}
