import com.azure.core.util.Context;

/** Samples for Projects GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/Projects_Get.json
     */
    /**
     * Sample code: Projects_Get.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void projectsGet(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.projects().getByResourceGroupWithResponse("rg1", "{projectName}", Context.NONE);
    }
}
