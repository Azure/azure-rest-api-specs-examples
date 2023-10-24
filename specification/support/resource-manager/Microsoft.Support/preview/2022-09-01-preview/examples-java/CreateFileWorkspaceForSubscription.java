/** Samples for FileWorkspaces Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/CreateFileWorkspaceForSubscription.json
     */
    /**
     * Sample code: Create a file workspace.
     *
     * @param manager Entry point to SupportManager.
     */
    public static void createAFileWorkspace(com.azure.resourcemanager.support.SupportManager manager) {
        manager.fileWorkspaces().createWithResponse("testworkspace", com.azure.core.util.Context.NONE);
    }
}
