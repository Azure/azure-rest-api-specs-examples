
import com.azure.resourcemanager.devcenter.models.CheckScopedNameAvailabilityRequest;

/**
 * Samples for CheckScopedNameAvailability Execute.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/
     * CheckScopedNameAvailability_ProjectCatalog.json
     */
    /**
     * Sample code: ProjectCatalogNameAvailability.
     * 
     * @param manager Entry point to DevCenterManager.
     */
    public static void projectCatalogNameAvailability(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.checkScopedNameAvailabilities().executeWithResponse(new CheckScopedNameAvailabilityRequest()
            .withName("name1").withType("Microsoft.DevCenter/projects/catalogs").withScope(
                "/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/projects/DevProject"),
            com.azure.core.util.Context.NONE);
    }
}
