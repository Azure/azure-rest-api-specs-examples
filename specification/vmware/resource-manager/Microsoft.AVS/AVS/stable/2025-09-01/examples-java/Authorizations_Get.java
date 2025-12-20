
/**
 * Samples for Authorizations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/Authorizations_Get.json
     */
    /**
     * Sample code: Authorizations_Get.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void authorizationsGet(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.authorizations().getWithResponse("group1", "cloud1", "authorization1",
            com.azure.core.util.Context.NONE);
    }
}
