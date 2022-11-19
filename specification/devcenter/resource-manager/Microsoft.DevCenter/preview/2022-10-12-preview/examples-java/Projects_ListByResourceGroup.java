import com.azure.core.util.Context;

/** Samples for Projects ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-10-12-preview/examples/Projects_ListByResourceGroup.json
     */
    /**
     * Sample code: Projects_ListByResourceGroup.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void projectsListByResourceGroup(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.projects().listByResourceGroup("rg1", null, Context.NONE);
    }
}
