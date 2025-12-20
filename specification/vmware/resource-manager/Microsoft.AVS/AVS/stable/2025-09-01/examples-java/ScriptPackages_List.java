
/**
 * Samples for ScriptPackages List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ScriptPackages_List.json
     */
    /**
     * Sample code: ScriptPackages_List.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void scriptPackagesList(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.scriptPackages().list("group1", "cloud1", com.azure.core.util.Context.NONE);
    }
}
