import com.azure.core.util.Context;

/** Samples for Workspaces ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/workspaces/Workspaces_ListByResourceGroup.json
     */
    /**
     * Sample code: Get workspaces by resource group.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void getWorkspacesByResourceGroup(
        com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.workspaces().listByResourceGroup("testRG", Context.NONE);
    }
}
