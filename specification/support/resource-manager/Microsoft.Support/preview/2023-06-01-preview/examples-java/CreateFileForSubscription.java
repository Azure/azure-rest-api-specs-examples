
/**
 * Samples for Files Create.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/
     * CreateFileForSubscription.json
     */
    /**
     * Sample code: Create a subscription scoped file.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void createASubscriptionScopedFile(com.azure.resourcemanager.support.SupportManager manager) {
        manager.files().define("test.txt").withExistingFileWorkspace("testworkspace").withChunkSize(41423)
            .withFileSize(41423).withNumberOfChunks(1).create();
    }
}
