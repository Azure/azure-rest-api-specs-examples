
/**
 * Samples for FileWorkspacesNoSubscription Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/CreateFileWorkspace.json
     */
    /**
     * Sample code: Create a file workspace.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void createAFileWorkspace(com.azure.resourcemanager.support.SupportManager manager) {
        manager.fileWorkspacesNoSubscriptions().createWithResponse("testworkspace", com.azure.core.util.Context.NONE);
    }
}
