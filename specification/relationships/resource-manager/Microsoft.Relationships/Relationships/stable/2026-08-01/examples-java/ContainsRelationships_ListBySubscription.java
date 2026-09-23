
/**
 * Samples for ContainsRelationships List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01/ContainsRelationships_ListBySubscription.json
     */
    /**
     * Sample code: ContainsRelationships_ListBySubscription.
     * 
     * @param manager Entry point to RelationshipsManager.
     */
    public static void
        containsRelationshipsListBySubscription(com.azure.resourcemanager.relationships.RelationshipsManager manager) {
        manager.containsRelationships().list(null, com.azure.core.util.Context.NONE);
    }
}
