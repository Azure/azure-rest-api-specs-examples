/** Samples for Projects ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/Projects_ListByResourceGroup.json
     */
    /**
     * Sample code: Projects_ListByResourceGroup.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void projectsListByResourceGroup(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.projects().listByResourceGroup("rg1", null, com.azure.core.util.Context.NONE);
    }
}
