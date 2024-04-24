
/**
 * Samples for FileWorkspaces Create.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/
     * CreateFileWorkspaceForSubscription.json
     */
    /**
     * Sample code: Create a file workspace for a subscription.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void createAFileWorkspaceForASubscription(com.azure.resourcemanager.support.SupportManager manager) {
        manager.fileWorkspaces().createWithResponse("testworkspace", com.azure.core.util.Context.NONE);
    }
}
