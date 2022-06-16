import com.azure.core.util.Context;

/** Samples for EntitiesRelations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/entities/relations/GetAllEntityRelations.json
     */
    /**
     * Sample code: Get all relations of an entity.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllRelationsOfAnEntity(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .entitiesRelations()
            .list("myRg", "myWorkspace", "afbd324f-6c48-459c-8710-8d1e1cd03812", null, null, null, null, Context.NONE);
    }
}
