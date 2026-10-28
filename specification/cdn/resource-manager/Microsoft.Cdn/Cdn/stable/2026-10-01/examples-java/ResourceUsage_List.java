
/**
 * Samples for ResourceUsage List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/ResourceUsage_List.json
     */
    /**
     * Sample code: ResourceUsage_List.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void resourceUsageList(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getResourceUsages().list(com.azure.core.util.Context.NONE);
    }
}
