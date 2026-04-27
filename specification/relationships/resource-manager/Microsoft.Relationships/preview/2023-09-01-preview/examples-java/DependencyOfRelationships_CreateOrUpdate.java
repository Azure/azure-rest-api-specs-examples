
import com.azure.resourcemanager.relationships.models.DependencyOfRelationshipProperties;

/**
 * Samples for DependencyOfRelationships CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-09-01-preview/DependencyOfRelationships_CreateOrUpdate.json
     */
    /**
     * Sample code: DependencyOfRelationships_CreateOrUpdate.
     * 
     * @param manager Entry point to RelationshipsManager.
     */
    public static void
        dependencyOfRelationshipsCreateOrUpdate(com.azure.resourcemanager.relationships.RelationshipsManager manager) {
        manager.dependencyOfRelationships().define("relationshipOne").withExistingResourceUri(
            "subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/resourceGroups/testrg/providers/Microsoft.DocumentDb/databaseAccounts/test-db-account")
            .withProperties(new DependencyOfRelationshipProperties().withTargetId(
                "/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/resourceGroups/testrg123/providers/Microsoft.Web/staticSites/test-site")
                .withTargetTenant("72f988bf-86f1-41af-91ab-2d7cd011db47"))
            .create();
    }
}
