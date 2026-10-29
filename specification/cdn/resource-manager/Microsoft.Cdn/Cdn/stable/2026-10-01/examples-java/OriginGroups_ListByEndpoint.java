
/**
 * Samples for OriginGroups ListByEndpoint.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/OriginGroups_ListByEndpoint.json
     */
    /**
     * Sample code: OriginsGroups_ListByEndpoint.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void originsGroupsListByEndpoint(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getOriginGroups().listByEndpoint("RG", "profile1", "endpoint1",
            com.azure.core.util.Context.NONE);
    }
}
