import com.azure.core.util.Context;

/** Samples for Workspaces GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/workspaces/Workspaces_Get.json
     */
    /**
     * Sample code: Get workspace.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void getWorkspace(com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.workspaces().getByResourceGroupWithResponse("testRG", "workspace1", Context.NONE);
    }
}
