
/**
 * Samples for StorageClass Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-03-01/StorageClass_Delete.json
     */
    /**
     * Sample code: StorageClass_Delete_0.
     * 
     * @param manager Entry point to ContainerOrchestratorRuntimeManager.
     */
    public static void storageClassDelete0(
        com.azure.resourcemanager.containerorchestratorruntime.ContainerOrchestratorRuntimeManager manager) {
        manager.storageClass().delete(
            "subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/example/providers/Microsoft.Kubernetes/connectedClusters/cluster1",
            "testrwx", com.azure.core.util.Context.NONE);
    }
}
