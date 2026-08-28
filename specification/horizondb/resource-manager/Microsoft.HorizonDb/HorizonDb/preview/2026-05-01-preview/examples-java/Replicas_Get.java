
/**
 * Samples for HorizonDbReplicas Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Replicas_Get.json
     */
    /**
     * Sample code: Get a HorizonDB replica.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void getAHorizonDBReplica(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.horizonDbReplicas().getWithResponse("exampleresourcegroup", "examplecluster", "examplepool",
            "examplereplica", com.azure.core.util.Context.NONE);
    }
}
