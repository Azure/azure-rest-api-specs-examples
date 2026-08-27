
/**
 * Samples for HorizonDbAdministrators Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Administrators_Get.json
     */
    /**
     * Sample code: Get a HorizonDB administrator.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void getAHorizonDBAdministrator(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.horizonDbAdministrators().getWithResponse("exampleresourcegroup", "examplecluster",
            "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", com.azure.core.util.Context.NONE);
    }
}
