import com.azure.resourcemanager.healthcareapis.models.WorkspaceProperties;

/** Samples for Workspaces CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/workspaces/Workspaces_Create.json
     */
    /**
     * Sample code: Create or update a workspace.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void createOrUpdateAWorkspace(
        com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager
            .workspaces()
            .define("workspace1")
            .withExistingResourceGroup("testRG")
            .withRegion("westus")
            .withProperties(new WorkspaceProperties())
            .create();
    }
}
