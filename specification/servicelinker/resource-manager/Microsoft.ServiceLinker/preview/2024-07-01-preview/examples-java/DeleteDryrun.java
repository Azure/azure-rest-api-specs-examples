
/**
 * Samples for LinkersOperation DeleteDryrun.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/
     * DeleteDryrun.json
     */
    /**
     * Sample code: DeleteDryrun.
     * 
     * @param manager Entry point to ServiceLinkerManager.
     */
    public static void deleteDryrun(com.azure.resourcemanager.servicelinker.ServiceLinkerManager manager) {
        manager.linkersOperations().deleteDryrunWithResponse(
            "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app",
            "dryrunName", com.azure.core.util.Context.NONE);
    }
}
