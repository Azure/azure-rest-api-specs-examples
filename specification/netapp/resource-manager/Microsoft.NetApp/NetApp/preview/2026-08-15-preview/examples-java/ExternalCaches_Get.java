
/**
 * Samples for ExternalCaches Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-15-preview/ExternalCaches_Get.json
     */
    /**
     * Sample code: ExternalCaches_Get.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void externalCachesGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.externalCaches().getWithResponse("contosoNetAppRg", "contosoExternalCluster", "contosoExternalSvm",
            "contosoCache_AB12", com.azure.core.util.Context.NONE);
    }
}
