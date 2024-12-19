
/**
 * Samples for Workspaces ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/workspaces/
     * Workspaces_ListByResourceGroup.json
     */
    /**
     * Sample code: Get workspaces by resource group.
     * 
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void
        getWorkspacesByResourceGroup(com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.workspaces().listByResourceGroup("testRG", com.azure.core.util.Context.NONE);
    }
}
