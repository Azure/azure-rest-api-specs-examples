/** Samples for Files Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/CreateFileForSubscription.json
     */
    /**
     * Sample code: Create a file workspace.
     *
     * @param manager Entry point to SupportManager.
     */
    public static void createAFileWorkspace(com.azure.resourcemanager.support.SupportManager manager) {
        manager
            .files()
            .define("test.txt")
            .withExistingFileWorkspace("testworkspace")
            .withChunkSize(41423.0F)
            .withFileSize(41423.0F)
            .withNumberOfChunks(1.0F)
            .create();
    }
}
