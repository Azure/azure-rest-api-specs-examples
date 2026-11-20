
/**
 * Samples for AccessPointResources List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-02-preview/AccessPointResources_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: AccessPointResources_List_MaximumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void
        accessPointResourcesListMaximumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.accessPointResources().list("rgconfluent", "myOrganization", "env-abc123", "gw-def456",
            com.azure.core.util.Context.NONE);
    }
}
