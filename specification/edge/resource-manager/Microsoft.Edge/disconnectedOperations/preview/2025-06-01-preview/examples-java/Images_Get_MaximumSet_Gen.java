
/**
 * Samples for Images Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/Images_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Images_Get.
     * 
     * @param manager Entry point to DisconnectedOperationsManager.
     */
    public static void
        imagesGet(com.azure.resourcemanager.disconnectedoperations.DisconnectedOperationsManager manager) {
        manager.images().getWithResponse("rgdisconnectedoperations", "bT62l-KS7g1-uh", "2P6",
            com.azure.core.util.Context.NONE);
    }
}
