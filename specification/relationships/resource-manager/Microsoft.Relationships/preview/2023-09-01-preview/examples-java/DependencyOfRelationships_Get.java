
/**
 * Samples for DependencyOfRelationships Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-09-01-preview/DependencyOfRelationships_Get.json
     */
    /**
     * Sample code: DependencyOfRelationships_Get.
     * 
     * @param manager Entry point to RelationshipsManager.
     */
    public static void
        dependencyOfRelationshipsGet(com.azure.resourcemanager.relationships.RelationshipsManager manager) {
        manager.dependencyOfRelationships().getWithResponse(
            "subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/resourceGroups/testrg/providers/Microsoft.DocumentDb/databaseAccounts/test-db-account",
            "relationshipOne", com.azure.core.util.Context.NONE);
    }
}
