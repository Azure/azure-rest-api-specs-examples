
/**
 * Samples for UnifiedResilienceItems List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/UnifiedResilienceItems_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: UnifiedResilienceItems_List_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void unifiedResilienceItemsListMaximumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.unifiedResilienceItems().list("zldmpkvqzifygkqau", "xntbyoswztnmvitj", 69,
            com.azure.core.util.Context.NONE);
    }
}
