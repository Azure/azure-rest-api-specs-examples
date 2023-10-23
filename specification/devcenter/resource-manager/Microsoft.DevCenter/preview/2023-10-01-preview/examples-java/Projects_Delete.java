/** Samples for Projects Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/Projects_Delete.json
     */
    /**
     * Sample code: Projects_Delete.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void projectsDelete(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.projects().delete("rg1", "DevProject", com.azure.core.util.Context.NONE);
    }
}
