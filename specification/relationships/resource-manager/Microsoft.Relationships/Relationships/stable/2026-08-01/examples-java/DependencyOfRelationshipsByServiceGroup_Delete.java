
/**
 * Samples for DependencyOfRelationshipsByServiceGroup Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01/DependencyOfRelationshipsByServiceGroup_Delete.json
     */
    /**
     * Sample code: DependencyOfRelationshipsByServiceGroup_Delete.
     * 
     * @param manager Entry point to RelationshipsManager.
     */
    public static void dependencyOfRelationshipsByServiceGroupDelete(
        com.azure.resourcemanager.relationships.RelationshipsManager manager) {
        manager.dependencyOfRelationshipsByServiceGroups().delete("myServiceGroup", "relationshipOne",
            com.azure.core.util.Context.NONE);
    }
}
