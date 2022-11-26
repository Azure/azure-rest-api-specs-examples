import com.azure.core.util.Context;

/** Samples for Projects List. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-11-11-preview/examples/Projects_ListBySubscription.json
     */
    /**
     * Sample code: Projects_ListBySubscription.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void projectsListBySubscription(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.projects().list(null, Context.NONE);
    }
}
