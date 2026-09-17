
import com.azure.resourcemanager.securityinsights.models.EntityQueryTemplateKind;

/**
 * Samples for EntityQueryTemplates List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/entityQueryTemplates/GetEntityQueryTemplates.json
     */
    /**
     * Sample code: Get all entity query templates.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getAllEntityQueryTemplates(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.entityQueryTemplates().list("myRg", "myWorkspace", EntityQueryTemplateKind.ACTIVITY,
            com.azure.core.util.Context.NONE);
    }
}
