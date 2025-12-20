
/**
 * Samples for ScriptPackages Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ScriptPackages_Get.json
     */
    /**
     * Sample code: ScriptPackages_Get.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void scriptPackagesGet(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.scriptPackages().getWithResponse("group1", "cloud1", "Microsoft.AVS.Management@3.0.48",
            com.azure.core.util.Context.NONE);
    }
}
