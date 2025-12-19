
/**
 * Samples for CloudLinks List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/CloudLinks_List.json
     */
    /**
     * Sample code: CloudLinks_List.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void cloudLinksList(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.cloudLinks().list("group1", "cloud1", com.azure.core.util.Context.NONE);
    }
}
