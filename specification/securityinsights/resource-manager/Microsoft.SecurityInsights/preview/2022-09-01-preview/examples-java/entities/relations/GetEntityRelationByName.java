import com.azure.core.util.Context;

/** Samples for EntityRelations GetRelation. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/entities/relations/GetEntityRelationByName.json
     */
    /**
     * Sample code: Get an entity relation.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAnEntityRelation(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .entityRelations()
            .getRelationWithResponse(
                "myRg",
                "myWorkspace",
                "afbd324f-6c48-459c-8710-8d1e1cd03812",
                "4bb36b7b-26ff-4d1c-9cbe-0d8ab3da0014",
                Context.NONE);
    }
}
