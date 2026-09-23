
/**
 * Samples for DependencyOfRelationshipsByServiceGroup Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01/DependencyOfRelationshipsByServiceGroup_Get.json
     */
    /**
     * Sample code: DependencyOfRelationshipsByServiceGroup_Get.
     * 
     * @param manager Entry point to RelationshipsManager.
     */
    public static void dependencyOfRelationshipsByServiceGroupGet(
        com.azure.resourcemanager.relationships.RelationshipsManager manager) {
        manager.dependencyOfRelationshipsByServiceGroups().getWithResponse("myServiceGroup", "relationshipOne",
            com.azure.core.util.Context.NONE);
    }
}
