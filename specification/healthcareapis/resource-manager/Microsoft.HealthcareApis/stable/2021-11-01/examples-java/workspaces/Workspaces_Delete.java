import com.azure.core.util.Context;

/** Samples for Workspaces Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/workspaces/Workspaces_Delete.json
     */
    /**
     * Sample code: Delete a workspace.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void deleteAWorkspace(com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.workspaces().delete("testRG", "workspace1", Context.NONE);
    }
}
