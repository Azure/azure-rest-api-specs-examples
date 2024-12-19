
/**
 * Samples for Files Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/CreateFileForSubscription.
     * json
     */
    /**
     * Sample code: Create a file under a subscription workspace.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void
        createAFileUnderASubscriptionWorkspace(com.azure.resourcemanager.support.SupportManager manager) {
        manager.files().define("test.txt").withExistingFileWorkspace("testworkspace").withChunkSize(41423)
            .withFileSize(41423).withNumberOfChunks(1).create();
    }
}
