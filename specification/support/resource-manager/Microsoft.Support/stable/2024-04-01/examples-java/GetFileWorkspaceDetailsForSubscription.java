
/**
 * Samples for FileWorkspaces Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/
     * GetFileWorkspaceDetailsForSubscription.json
     */
    /**
     * Sample code: Get details of a subscription file workspace.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void
        getDetailsOfASubscriptionFileWorkspace(com.azure.resourcemanager.support.SupportManager manager) {
        manager.fileWorkspaces().getWithResponse("testworkspace", com.azure.core.util.Context.NONE);
    }
}
