
/**
 * Samples for ContainsRelationships ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01/ContainsRelationships_ListByResourceGroup.json
     */
    /**
     * Sample code: ContainsRelationships_ListByResourceGroup.
     * 
     * @param manager Entry point to RelationshipsManager.
     */
    public static void
        containsRelationshipsListByResourceGroup(com.azure.resourcemanager.relationships.RelationshipsManager manager) {
        manager.containsRelationships().listByResourceGroup("testrg", null, com.azure.core.util.Context.NONE);
    }
}
