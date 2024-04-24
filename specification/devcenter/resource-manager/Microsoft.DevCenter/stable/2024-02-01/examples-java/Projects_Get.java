
/**
 * Samples for Projects GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/Projects_Get.json
     */
    /**
     * Sample code: Projects_Get.
     * 
     * @param manager Entry point to DevCenterManager.
     */
    public static void projectsGet(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.projects().getByResourceGroupWithResponse("rg1", "DevProject", com.azure.core.util.Context.NONE);
    }
}
