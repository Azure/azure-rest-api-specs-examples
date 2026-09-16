
import com.azure.resourcemanager.securityinsights.models.EntityItemQueryKind;

/**
 * Samples for Entities Queries.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/entities/GetQueries.json
     */
    /**
     * Sample code: Get Entity Query.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getEntityQuery(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.entities().queries("myRg", "myWorkspace", "e1d3d618-e11f-478b-98e3-bb381539a8e1",
            EntityItemQueryKind.INSIGHT, com.azure.core.util.Context.NONE);
    }
}
