
/**
 * Samples for HorizonDbPrivateLinkResources Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/PrivateLinkResources_Get.json
     */
    /**
     * Sample code: Gets a private link resource for HorizonDB.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        getsAPrivateLinkResourceForHorizonDB(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.horizonDbPrivateLinkResources().getWithResponse("exampleresourcegroup", "examplecluster", "default",
            com.azure.core.util.Context.NONE);
    }
}
