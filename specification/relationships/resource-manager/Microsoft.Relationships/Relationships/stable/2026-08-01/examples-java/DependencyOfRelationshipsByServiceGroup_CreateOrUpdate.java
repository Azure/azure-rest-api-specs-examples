
import com.azure.resourcemanager.relationships.fluent.models.DependencyOfRelationshipInner;
import com.azure.resourcemanager.relationships.models.DependencyOfRelationshipProperties;

/**
 * Samples for DependencyOfRelationshipsByServiceGroup CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01/DependencyOfRelationshipsByServiceGroup_CreateOrUpdate.json
     */
    /**
     * Sample code: DependencyOfRelationshipsByServiceGroup_CreateOrUpdate.
     * 
     * @param manager Entry point to RelationshipsManager.
     */
    public static void dependencyOfRelationshipsByServiceGroupCreateOrUpdate(
        com.azure.resourcemanager.relationships.RelationshipsManager manager) {
        manager.dependencyOfRelationshipsByServiceGroups().createOrUpdate("myServiceGroup", "relationshipOne",
            new DependencyOfRelationshipInner().withProperties(new DependencyOfRelationshipProperties().withTargetId(
                "/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/resourceGroups/testrg123/providers/Microsoft.Web/staticSites/test-site")
                .withTargetTenant("72f988bf-86f1-41af-91ab-2d7cd011db47")),
            com.azure.core.util.Context.NONE);
    }
}
