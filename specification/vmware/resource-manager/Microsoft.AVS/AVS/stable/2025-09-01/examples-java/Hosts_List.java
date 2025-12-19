
/**
 * Samples for Hosts List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/Hosts_List.json
     */
    /**
     * Sample code: Hosts_List.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void hostsList(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.hosts().list("group1", "cloud1", "cluster1", com.azure.core.util.Context.NONE);
    }
}
