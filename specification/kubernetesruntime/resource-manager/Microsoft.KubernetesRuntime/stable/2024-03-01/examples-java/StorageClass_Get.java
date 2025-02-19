
/**
 * Samples for StorageClass Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-03-01/StorageClass_Get.json
     */
    /**
     * Sample code: StorageClass_Get_0.
     * 
     * @param manager Entry point to ContainerOrchestratorRuntimeManager.
     */
    public static void storageClassGet0(
        com.azure.resourcemanager.containerorchestratorruntime.ContainerOrchestratorRuntimeManager manager) {
        manager.storageClass().getWithResponse(
            "subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/example/providers/Microsoft.Kubernetes/connectedClusters/cluster1",
            "testrwx", com.azure.core.util.Context.NONE);
    }
}
