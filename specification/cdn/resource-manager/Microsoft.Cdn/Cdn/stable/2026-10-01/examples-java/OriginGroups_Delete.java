
/**
 * Samples for OriginGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/OriginGroups_Delete.json
     */
    /**
     * Sample code: OriginGroups_Delete.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void originGroupsDelete(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getOriginGroups().delete("RG", "profile1", "endpoint1", "originGroup1",
            com.azure.core.util.Context.NONE);
    }
}
