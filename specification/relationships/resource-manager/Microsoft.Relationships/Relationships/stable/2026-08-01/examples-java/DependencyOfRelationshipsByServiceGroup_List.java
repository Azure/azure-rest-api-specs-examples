
/**
 * Samples for DependencyOfRelationshipsByServiceGroup List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01/DependencyOfRelationshipsByServiceGroup_List.json
     */
    /**
     * Sample code: DependencyOfRelationshipsByServiceGroup_List.
     * 
     * @param manager Entry point to RelationshipsManager.
     */
    public static void dependencyOfRelationshipsByServiceGroupList(
        com.azure.resourcemanager.relationships.RelationshipsManager manager) {
        manager.dependencyOfRelationshipsByServiceGroups().list("myServiceGroup", com.azure.core.util.Context.NONE);
    }
}
