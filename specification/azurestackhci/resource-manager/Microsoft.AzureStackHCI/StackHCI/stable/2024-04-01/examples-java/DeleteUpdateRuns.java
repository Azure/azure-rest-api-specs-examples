
/**
 * Samples for UpdateRuns Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/
     * DeleteUpdateRuns.json
     */
    /**
     * Sample code: Delete an Update.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void deleteAnUpdate(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.updateRuns().delete("testrg", "testcluster", "Microsoft4.2203.2.32",
            "23b779ba-0d52-4a80-8571-45ca74664ec3", com.azure.core.util.Context.NONE);
    }
}
